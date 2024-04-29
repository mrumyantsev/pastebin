package endpoint

import (
	"github.com/gin-gonic/gin"
	docs "github.com/mrumyantsev/pastebin/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (e *Endpoint) InitRoutes(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/sign-up", e.Authorization.SignUp)
			auth.POST("/sign-in", e.Authorization.SignIn)
			auth.GET("/validate", e.Authorization.Validate)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
