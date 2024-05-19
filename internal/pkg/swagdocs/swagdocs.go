package swagdocs

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mrumyantsev/pastebin/docs"
)

// InitRoutes initiates an additional HTTP route for getting Swagger
// documentation.
func InitRoutes(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
