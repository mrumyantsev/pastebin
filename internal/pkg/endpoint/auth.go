package endpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
)

type AuthorizationEndpoint struct {
	config *config.Config
}

func NewAuthorizationEndpoint(cfg *config.Config) *AuthorizationEndpoint {
	return &AuthorizationEndpoint{config: cfg}
}

func (e *AuthorizationEndpoint) SignUp(ctx *gin.Context) {
	fmt.Println("sign up")
}

func (e *AuthorizationEndpoint) SignIn(ctx *gin.Context) {
	fmt.Println("sign in")
}

func (e *AuthorizationEndpoint) Validate(ctx *gin.Context) {
	fmt.Println("validate")
}
