package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type NoteContent struct {
	NoteID      uuid.UUID `gorm:"type:char(36);not null;primary_key"`
	Content     string    `gorm:"type:text;not null"`
	ExpiredTime sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NoteResponse struct {
	NoteID      uuid.UUID
	Content     string
	ExpiredTime *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
