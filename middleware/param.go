package middleware

import (
	"github.com/KashEight/not/consts"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetUUIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("uuid")

		uuid.MustParse(v)

		ctx.Set(consts.UUIDKeyName, v)

		ctx.Next()
	}
}
