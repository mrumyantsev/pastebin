package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin/pkg/lib/errlib"
	"github.com/rs/zerolog/log"
)

const (
	message = "message"
)

func Fail(ctx *gin.Context, err error, msg string) {
	if errlib.IsCustom(err) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{message: err.Error()})
	} else {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	log.Error().Err(err).Msg(msg)
}

func Success(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{message: msg})

	log.Info().Msg(msg)
}

func Object(ctx *gin.Context, obj any, msg string) {
	ctx.JSON(http.StatusOK, obj)

	log.Info().Msg(msg)
}
