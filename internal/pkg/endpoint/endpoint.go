package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
)

type Authorization interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	Validate(ctx *gin.Context)
}

type Endpoint struct {
	Authorization Authorization
}

func New(cfg *config.Config) *Endpoint {
	return &Endpoint{
		Authorization: NewAuthorizationEndpoint(cfg),
	}
}
