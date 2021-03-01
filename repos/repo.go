package repos

import "gorm.io/gorm"

type Repo interface {
	NoteRepository
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repo {
	r := &repo{
		db: db,
	}
	return r
}
