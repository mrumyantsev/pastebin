package v1

import "github.com/gin-gonic/gin"

func (e *UserV1HttpEndpoint) InitRoutes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", e.CreateUser)
			users.GET("/", e.GetUsers)
			users.GET("/:id", e.GetUser)
			users.PATCH("/:id", e.UpdateUser)
			users.DELETE("/:id", e.DeleteUser)

			users.GET("/is-exists/:username", e.IsUserExists)
			users.GET("/is-email-exists/:email", e.IsEmailExists)
			users.GET("/count", e.UserCount)
		}
	}
}
