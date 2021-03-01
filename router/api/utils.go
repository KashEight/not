package api

import (
	"github.com/KashEight/not/consts"
	"github.com/KashEight/not/models"
	"github.com/gin-gonic/gin"
)

func getUUIDKey(ctx *gin.Context) string {
	uuid, _ := ctx.Get(consts.UUIDKeyName)
	return uuid.(string)
}

func validateNoteContent(nc *models.NoteContent) *models.NoteResponse {
	nr := &models.NoteResponse{
		NoteID:    nc.NoteID,
		Content:   nc.Content,
		CreatedAt: nc.CreatedAt,
		UpdatedAt: nc.UpdatedAt,
	}

	if nc.ExpiredTime.Valid {
		nr.ExpiredTime = &nc.ExpiredTime.Time
	}

	return nr
}

func validateNoteContentSlice(nc []models.NoteContent) []models.NoteResponse {
	nrs := make([]models.NoteResponse, 0)

	for _, ncp := range nc {
		nr := validateNoteContent(&ncp)
		nrs = append(nrs, *nr)
	}

	return nrs
}

func validateNotePostData(pd models.NotePostData) (*models.NoteContent, error) {
	nc, err := pd.ConvertToNoteContent()

	if err != nil {
		return nil, err
	}

	return nc, nil
}
