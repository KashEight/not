package repos

import (
	"github.com/KashEight/not/models"
	"gorm.io/gorm"
)

type NoteRepository interface {
	GetAllNote() ([]models.NoteContent, error)
	GetNoteByUUID(uuid string) (*models.NoteContent, error)
	CreateNote(nc *models.NoteContent) error
	UpdateNote(uuid string, nc *models.NoteContent) error
	DeleteNote(uuid string) error
}

func (r *repo) GetAllNote() ([]models.NoteContent, error) {
	notes := make([]models.NoteContent, 0)
	tx := r.db.Find(&notes)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repo) GetNoteByUUID(uuid string) (*models.NoteContent, error) {
	note := &models.NoteContent{}
	tx := r.db.Where("note_id = ?", uuid).First(note)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return note, nil
}

func (r *repo) CreateNote(nc *models.NoteContent) error {
	tx := r.db.Create(nc)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateNote(uuid string, nc *models.NoteContent) error {
	tx := r.db.Model(&nc).Where("note_id = ?", uuid).Updates(nc)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteNote(uuid string) error {
	tx := r.db.Where("note_id = ?", uuid).Delete(&models.NoteContent{})

	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if err := tx.Error; err != nil {
		return err
	}

	return nil
}
