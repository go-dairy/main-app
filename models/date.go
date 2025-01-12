package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrDateExists   = errors.New("date exists")
	ErrDateNotFound = errors.New("date not found")
)

type Date struct {
	Date_Uuid uuid.UUID
	Date      time.Time
}
