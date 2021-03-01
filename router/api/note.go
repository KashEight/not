package api

import (
	"errors"
	"github.com/KashEight/not/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *handler) getAllNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		notesRaw, err := h.repo.GetAllNote()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			notes := validateNoteContentSlice(notesRaw)
			ctx.JSON(http.StatusOK, notes)
		}
	}
}

func (h *handler) createNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &models.PostDataCreateNote{}

		if err := ctx.BindJSON(pd); err != nil {
			return
		}

		noteRaw, err := validateNotePostData(pd)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.repo.CreateNote(noteRaw); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			note := validateNoteContent(noteRaw)
			ctx.JSON(http.StatusCreated, note)
		}
	}
}

func (h *handler) getNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := getUUIDKey(ctx)

		noteRaw, err := h.repo.GetNoteByUUID(uuid)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(http.StatusBadRequest)
		} else if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			note := validateNoteContent(noteRaw)
			ctx.JSON(http.StatusOK, note)
		}
	}
}

func (h *handler) updateNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &models.PostDataUpdateNote{}

		if err := ctx.BindJSON(pd); err != nil {
			return
		}

		noteRaw, err := validateNotePostData(pd)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		uuid := getUUIDKey(ctx)

		if err := h.repo.UpdateNote(uuid, noteRaw); errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(http.StatusBadRequest)
		} else if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			note := validateNoteContent(noteRaw)
			ctx.JSON(http.StatusNoContent, note)
		}
	}
}

func (h *handler) deleteNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := getUUIDKey(ctx)

		if err := h.repo.DeleteNote(uuid); errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(http.StatusBadRequest)
		} else if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusOK)
		}
	}
}
