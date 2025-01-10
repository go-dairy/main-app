package godairy

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-dairy/main-app/internal/logger/sl"
	"github.com/go-dairy/main-app/internal/models"
	"github.com/google/uuid"
)

type TaskSaver interface {
	SaveTask(
		ctx context.Context,
		task models.Task,
	) (record_id uuid.UUID, err error)

	NewDate(
		ctx context.Context,
		date time.Time,
	) (dairyDate *models.Date, err error)

	DeleteTask(
		ctx context.Context,
		record_id uuid.UUID,
	) (err error)
}

type TaskProvider interface {
	TaskByUuid(
		ctx context.Context,
		record_id uuid.UUID,
	) (task *models.Task, err error)
	TaskByDateTime(
		ctx context.Context,
		date time.Time,
		time time.Time,
	) (tasks []*models.Task, err error)
	Date(
		ctx context.Context,
		date time.Time,
	) (dt *models.Date, err error)
}

type Godairy struct {
	log          *slog.Logger
	taskSaver    TaskSaver
	taskProvider TaskProvider
}

// New returns a new instance of the Godairy service.
func New(
	log *slog.Logger,
	taskSaver TaskSaver,
	taskProvider TaskProvider,
) *Godairy {
	return &Godairy{
		log:          log,
		taskSaver:    taskSaver,
		taskProvider: taskProvider,
	}
}

// NewTask adds a new task record into storage.
//
// If the task already exists, returns an error indicating the task exists.
// Logs the operation and any errors encountered during the process.
//
// Parameters:
//   ctx - The context for the request, used for cancellation and deadlines.
//   task - The task object containing details of the task to be saved.
//
// Returns:
//   record_id - The UUID of the newly saved task.
//   err - An error object if the operation fails or if the task already exists.

func (g *Godairy) NewTask(
	ctx context.Context,
	taskDateTime time.Time,
	taskBody string,
) (uuid.UUID, error) {
	const op = "internal.services.godairy.NewTask"
	log := g.log.With(slog.String("op", op), slog.Any("DateTime", taskDateTime), slog.String("Body", taskBody))

	log.Info("adding new task record")

	date, err := g.taskProvider.Date(ctx, taskDateTime)
	if err != nil {
		if errors.Is(err, models.ErrDateNotFound) {
			log.Info("date not found, creating it", sl.Err(err))

			if date, err = g.taskSaver.NewDate(ctx, taskDateTime); err != nil {
				log.Error("failed to save date", sl.Err(err))

				return uuid.Nil, fmt.Errorf("%s: %w", op, err)
			}
		} else {
			log.Error("failed to get date", sl.Err(err))

			return uuid.Nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	task := models.Task{
		Time:      taskDateTime,
		Task:      taskBody,
		Date_Uuid: date.Date_Uuid,
	}

	id, err := g.taskSaver.SaveTask(ctx, task)
	if err != nil {
		if errors.Is(err, models.ErrTaskExists) {
			log.Info("task exists", sl.Err(err))

			return uuid.Nil, fmt.Errorf("%s: %w", op, models.ErrTaskExists)
		}
		log.Error("failed to save task", sl.Err(err))

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("new task successfully added")

	return id, nil
}

// NewDate adds a new date record into storage.
//
// If the date already exists, returns an error indicating the date exists.
// Logs the operation and any errors encountered during the process.
//
// Parameters:
//   ctx - The context for controlling cancellation and timeouts.
//   date - The date to be added to storage.
//
// Returns:
//   dairyDate - The newly created date record.
//   err - An error if the date cannot be added or another error occurs.

func (g *Godairy) NewDate(
	ctx context.Context,
	date time.Time,
) (dairyDate *models.Date, err error) {
	const op = "internal.services.godairy.NewDate"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("date", date),
	)

	log.Info("adding new date record")

	dairyDate, err = g.taskSaver.NewDate(ctx, date)
	if errors.Is(err, models.ErrDateExists) {
		log.Error("date exists", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, models.ErrTaskExists)
	}
	if err != nil {
		log.Error("failed to save date", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return dairyDate, nil
}

// Date retrieves a date record from storage based on the provided date.
//
// If the date doesn't exist, returns an error indicating the date was not found.
// Returns the date record and nil error if successful.
//
// Parameters:
//   ctx - The context for controlling cancellation and timeouts.
//   date - The date to be retrieved.
//
// Returns:
//   dt - The date record associated with the provided date.
//   err - An error if the date cannot be found or another error occurs.

func (g *Godairy) Date(
	ctx context.Context,
	date time.Time,
) (dt *models.Date, err error) {
	const op = "internal.services.godairy.Date"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("date", date),
	)

	log.Info("finding date")

	dt, err = g.taskProvider.Date(ctx, date)
	if errors.Is(err, models.ErrDateNotFound) {
		log.Info("date not found", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err != nil {
		log.Error("failed to get date", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return dt, nil
}

// Task retrieves tasks based on the provided task details.
//
// If the Task_Uuid is not nil, it retrieves a task by its unique identifier.
// If the time is specified, it retrieves tasks that match the given time.
// Returns an error if the task cannot be found or another error occurs.
//
// Parameters:
//   ctx - The context for controlling cancellation and timeouts.
//   task - The task model containing the details for retrieving tasks.
//
// Returns:
//   tasks - A slice of task pointers that match the given criteria.
//   err - An error if no tasks are found or another error occurs.

func (g *Godairy) Task(
	ctx context.Context,
	task models.Task,
) (tasks []*models.Task, err error) {
	const op = "internal.services.godairy.Task"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("task", task),
	)

	result := make([]*models.Task, 0, 1)

	switch {
	case task.Task_Uuid != uuid.Nil:
		log.Info("getting task by Uuid")
		j, err := g.taskProvider.TaskByUuid(ctx, task.Task_Uuid)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		if j != nil {
			result = append(result, j)
		}

		return result, nil

	case task.Time != time.Time{}:
		log.Info("getting task by time")
		result, err = g.taskProvider.TaskByDateTime(ctx, time.Time{}, task.Time)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return result, nil

	default:
		return nil, fmt.Errorf("%s: %w", op, models.ErrTaskNotFound)
	}
}

// TaskByDateTime retrieves tasks from storage based on the provided date and
// time.
//
// If the time is not specified, it retrieves all tasks for the given date.
// Returns the tasks and nil error if successful.
//
// Parameters:
//
//	ctx - The context for controlling cancellation and timeouts.
//	date - The date to be retrieved.
//	time - The time to be retrieved.
//
// Returns:
//
//	tasks - A slice of task pointers that match the given criteria.
//	err - An error if no tasks are found or another error occurs.
func (g *Godairy) TaskByDateTime(
	ctx context.Context,
	date time.Time,
	time time.Time,
) (tasks []*models.Task, err error) {
	const op = "internal.services.godairy.TaskByDateTime"

	log := g.log.With(
		slog.String("op", op),
	)

	log.Info("retrieving tasks by datetime")

	tasks, err = g.taskProvider.TaskByDateTime(ctx, date, time)
	if err != nil {
		log.Error("failed to retrive tasks by datetime", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tasks, nil
}

// DeleteTask deletes the task with the provided task uuid from the storage.
//
// Parameters:
//
//	ctx - The context for controlling cancellation and timeouts.
//	task - The task to be deleted.
//
// Returns:
//
//	affected - The number of affected rows.
//	status - True if the task was successfully deleted, false otherwise.
//	err - An error if the task does not exist or another error occurs.
func (g *Godairy) DeleteTask(
	ctx context.Context,
	task_uuid uuid.UUID,
) (err error) {
	const op = "internal.services.godairy.DeleteTask"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("task_uuid", task_uuid),
	)

	log.Info("deleting task")

	err = g.taskSaver.DeleteTask(ctx, task_uuid)
	if err != nil {
		log.Error("failed to delete task", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
