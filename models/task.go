package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTaskExists    = errors.New("task exists")
	ErrTaskNotFound  = errors.New("task not found")
	ErrTaskScheduled = errors.New("task already scheduled for this time")
)

type Task struct {
	Task_Uuid uuid.UUID
	Time      time.Time
	Task      string
	Date_Uuid uuid.UUID
}
