package grpcserver

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-dairy/main-app/internal/logger/sl"
	"github.com/go-dairy/main-app/internal/models"
	godairyv1 "github.com/go-dairy/main-app/protos/gen/go/dairy"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	TimeRounding = time.Duration(5 * time.Minute)
)

type Godairy interface {
	NewTask(
		ctx context.Context,
		taskDateTime time.Time,
		taskBody string,
	) (uuid uuid.UUID, err error)
	NewDate(
		ctx context.Context,
		date time.Time,
	) (dairyDate *models.Date, err error)
	Date(
		ctx context.Context,
		date time.Time,
	) (dairyDate *models.Date, err error)
	Task(
		ctx context.Context,
		task models.Task,
	) (tasks []*models.Task, err error)
	DeleteTask(
		ctx context.Context,
		task_uuid uuid.UUID,
	) (err error)
}

type serverAPI struct {
	godairyv1.UnimplementedGodairyServer
	godairy Godairy
}

// Register registers the Godairy service with the gRPC server.
//
// Parameters:
//
//	gRPC - The gRPC server to which the service is registered.
//	godairy - The implementation of the Godairy service to be registered.
func Register(gRPC *grpc.Server, godairy Godairy) {
	godairyv1.RegisterGodairyServer(gRPC, &serverAPI{godairy: godairy})
}

// NewTask creates a new task in the godairy service.
//
// Parameters:
//
//	ctx - The context for controlling cancellation and timeouts.
//	req - The request containing the task to be created.
//
// Returns:
//
//	response - The response containing the uuid of the created task.
//	err - An error if the request is invalid, a task with the same uuid already exists, or another error occurs.
func (s *serverAPI) NewTask(ctx context.Context,
	req *godairyv1.NewTaskRequest) (*godairyv1.NewTaskResponse, error) {
	var err error
	if err = validator.New().Struct(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	dateTime := req.GetDatetime().AsTime()

	record_id, err := s.godairy.NewTask(ctx, dateTime, req.GetTask())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &godairyv1.NewTaskResponse{RecordUuid: record_id.String()}, nil
}

// NewDate creates a new date in the godairy service.
//
// Parameters:
//
//	ctx - The context for controlling cancellation and timeouts.
//	req - The request containing the date to be created.
//
// Returns:
//
//	response - The response containing the uuid of the created date.
//	err - An error if the request is invalid, a date with the same uuid already exists, or another error occurs.
func (s *serverAPI) NewDate(ctx context.Context,
	req *godairyv1.DateRequest) (*godairyv1.DateResponse,
	error) {

	if err := validator.New().Struct(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	dt := req.GetDate().AsTime()

	dairyDate, err := s.godairy.NewDate(ctx, dt)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot create date")
	}

	return &godairyv1.DateResponse{
		DateUuid: dairyDate.Date_Uuid.String(),
		Date:     timestamppb.New(dairyDate.Date),
	}, nil
}

// Date retrieves a date record from storage based on the provided date.
//
// Parameters:
//
//	ctx - The context for controlling cancellation and timeouts.
//	req - The request containing the date to be retrieved.
//
// Returns:
//
//	response - The response containing the uuid of the retrieved date.
//	err - An error if the request is invalid, a date with the same uuid already exists, or another error occurs.
func (s *serverAPI) Date(ctx context.Context,
	req *godairyv1.DateRequest) (*godairyv1.DateResponse, error) {

	date := req.GetDate().AsTime()

	dairyDate, err := s.godairy.Date(ctx, date)
	if err != nil {
		slog.Info("cannot get date", sl.Err(err))
		return nil, status.Error(codes.Internal, "cannot get date")
	}

	return &godairyv1.DateResponse{
		DateUuid: dairyDate.Date_Uuid.String(),
		Date:     timestamppb.New(dairyDate.Date),
	}, nil
}

// Task retrieves tasks based on the provided request.
//
// Validates the request, parses the task UUID and datetime, and retrieves tasks
// that match the specified criteria from the godairy service.
//
// Parameters:
//  ctx - The context for controlling cancellation and timeouts.
//  req - The request containing the task details to be retrieved.
//
// Returns:
//  response - The response containing the list of tasks that match the criteria.
//  err - An error if the request is invalid, the UUID or datetime cannot be
//        parsed, or tasks cannot be retrieved.

func (s *serverAPI) Task(ctx context.Context,
	req *godairyv1.TaskRequest) (*godairyv1.TaskResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		slog.Info("invalid request", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	u, err := uuid.Parse(req.GetRecordUuid())
	if err != nil {
		slog.Info("cannot parse uuid", sl.Err(err))
		return nil, status.Error(codes.Internal, "cannot parse uuid")
	}

	t := req.GetDatetime().AsTime()

	t = t.Truncate(TimeRounding)

	tasks, err := s.godairy.Task(ctx, models.Task{
		Task_Uuid: u,
		Task:      req.GetTask(),
		Time:      t,
	})
	if err != nil {
		slog.Error("cannot get tasks", sl.Err(err))
		return nil, status.Error(codes.Internal, "cannot get tasks")
	}

	j := make([]*godairyv1.Task, len(tasks))
	for i, task := range tasks {
		j[i] = &godairyv1.Task{
			RecordUuid: task.Task_Uuid.String(),
			Datetime:   timestamppb.New(task.Time),
			Task:       task.Task,
		}
	}

	return &godairyv1.TaskResponse{
		Tasks: j,
	}, nil
}

// DeleteTask deletes a task based on the provided request.
//
// Validates the request, parses the task UUID, and deletes the task from the godairy service.
//
// Parameters:
//  ctx - The context for controlling cancellation and timeouts.
//  req - The request containing the task UUID to be deleted.
//
// Returns:
//  response - The response containing the number of affected rows and a boolean indicating if the task was deleted.
//  err - An error if the request is invalid, the UUID cannot be parsed, or the task cannot be deleted.

func (s *serverAPI) DeleteTask(ctx context.Context,
	req *godairyv1.DeleteTaskRequest) (*emptypb.Empty, error) {
	if err := validator.New().Struct(req); err != nil {
		return &emptypb.Empty{}, nil
	}

	u, err := uuid.Parse(req.GetRecordUuid())
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "cannot parse uuid")
	}

	err = s.godairy.DeleteTask(ctx, u)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.NotFound, "cannot delete task")
	}

	return &emptypb.Empty{}, nil
}
