package models

import (
	"database/sql"
	"github.com/KashEight/not/utils"
	"github.com/google/uuid"
	"time"
)

type NotePostData interface {
	ConvertToNoteContent() (*NoteContent, error)
}

type PostDataCreateNote struct {
	Content     string     `json:"content"`
	ExpiredTime *time.Time `json:"expired_time" time_format:""` // Empty string because time_format should be RFC3339 format
}

func (pd *PostDataCreateNote) ConvertToNoteContent() (*NoteContent, error) {
	c := pd.Content
	et := pd.ExpiredTime

	if c == "" && et == nil {
		return nil, utils.ErrInvalidPostData
	}

	t := sql.NullTime{}

	if et != nil {
		t.Time = *et
		t.Valid = true
	} else {
		t.Valid = false
	}

	newUUID, _ := uuid.NewRandom()

	note := &NoteContent{
		NoteID:      newUUID,
		Content:     c,
		ExpiredTime: t,
	}

	return note, nil
}

type PostDataUpdateNote struct {
	Content     string     `json:"content"`
	ExpiredTime *time.Time `json:"expired_time" time_format:""` // Empty string because time_format should be RFC3339 format
}

func (pd *PostDataUpdateNote) ConvertToNoteContent() (*NoteContent, error) {
	c := pd.Content
	et := pd.ExpiredTime

	if c == "" && et == nil {
		return nil, utils.ErrInvalidPostData
	}

	t := sql.NullTime{}

	if et != nil {
		t.Time = *et
		t.Valid = true
	} else {
		t.Valid = false
	}

	note := &NoteContent{
		Content:     c,
		ExpiredTime: t,
	}

	return note, nil
}
