package pgsql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-dairy/main-app/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	DB *pgxpool.Pool
}

func New(dsn string) (*Storage, error) {
	const op = "storage.pgsql.New"

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) Close() error {
	s.DB.Close()

	return nil
}

func (s *Storage) SaveTask(
	ctx context.Context,
	task models.Task,
) (record_id uuid.UUID, err error) {
	const op = "storage.pgsql.SaveTask"
	query := `
INSERT INTO tasks (
	dairy_time,
	dairy_task,
	date_uuid
) VALUES (
	@dairy_time,
	@dairy_task,
	@date_uuid
) RETURNING task_uuid;`

	args := pgx.NamedArgs{
		"dairy_time": task.Time,
		"dairy_task": task.Task,
		"date_uuid":  task.Date_Uuid,
	}
	err = s.DB.QueryRow(ctx, query, args).Scan(&task.Task_Uuid)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == "idx_time" {
			return uuid.Nil, fmt.Errorf("%s: %w", op, models.ErrTaskScheduled)
		}
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return task.Task_Uuid, nil
}

func (s *Storage) DeleteTask(
	ctx context.Context,
	record_id uuid.UUID,
) (err error) {
	const op = "storage.pgsql.DeleteTask"
	query := `
DELETE FROM tasks
WHERE task_uuid = @task_uuid;
`

	args := pgx.NamedArgs{
		"task_uuid": record_id,
	}
	_, err = s.DB.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) TaskByUuid(
	ctx context.Context,
	record_id uuid.UUID,
) (task *models.Task, err error) {
	const op = "storage.pgsql.TaskById"
	query := `
SELECT 
	task_uuid,
	dairy_time,
	dairy_task,
	date_uuid
FROM tasks
WHERE task_uuid = @task_uuid;
`

	args := pgx.NamedArgs{
		"task_uuid": record_id,
	}
	task = &models.Task{}
	err = s.DB.QueryRow(ctx, query, args).Scan(&task.Task_Uuid, &task.Time, &task.Task, &task.Date_Uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return task, nil
}

func (s *Storage) TaskByDateTime(
	ctx context.Context,
	tdate time.Time,
	ttime time.Time,
) (tasks []*models.Task, err error) {
	const op = "storage.pgsql.TaskByDateTime"
	var query string
	var args pgx.NamedArgs

	if !tdate.IsZero() {
		dt, err := s.Date(ctx, tdate)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		query = `
SELECT 
	task_uuid,
	dairy_time,
	dairy_task,
	date_uuid
FROM tasks
WHERE date_uuid = @date_uuid;`

		args = pgx.NamedArgs{
			"date_uuid": dt.Date_Uuid,
		}

	}
	if !ttime.IsZero() {
		query = `
SELECT 
	task_uuid,
	dairy_time,
	dairy_task,
	date_uuid
FROM tasks
WHERE dairy_time = @dairy_time;`

		args = pgx.NamedArgs{
			"dairy_time": ttime.Format(time.TimeOnly),
		}
	}
	rows, err := s.DB.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		task := &models.Task{}
		err = rows.Scan(&task.Task_Uuid, &task.Time, &task.Task, &task.Date_Uuid)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Storage) Date(
	ctx context.Context,
	date time.Time,
) (dt *models.Date, err error) {
	const op = "storage.pgsql.Date"
	query := `
SELECT 
	date_uuid,
	dairy_date
FROM dates
WHERE dairy_date = @date;	
`

	args := pgx.NamedArgs{
		"date": date.Format(time.DateOnly),
	}
	dt = &models.Date{}
	err = s.DB.QueryRow(ctx, query, args).Scan(&dt.Date_Uuid, &dt.Date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, models.ErrDateNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return dt, nil
}

func (s *Storage) NewDate(
	ctx context.Context,
	date time.Time,
) (dairyDate *models.Date, err error) {
	const op = "storage.pgsql.NewDate"

	query := `
INSERT INTO dates (
	dairy_date
) VALUES (
	@date
) RETURNING date_uuid;
 `

	args := pgx.NamedArgs{
		"date": date.Format(time.DateOnly),
	}
	dairyDate = &models.Date{Date: date}
	err = s.DB.QueryRow(ctx, query, args).Scan(&dairyDate.Date_Uuid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("%s: %w", op, models.ErrDateNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return dairyDate, nil
}
