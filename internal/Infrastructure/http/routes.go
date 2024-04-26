package http

import (
	"fmt"
	"github.com/andreluizmicro/desafio-go/internal/Infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	userController *controller.UserController,
	port string,
) {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		c.Writer.Header().Set("Expires", "0")                                         // Proxies.
		c.Next()
	})

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		users.POST("/", userController.Create)

	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	}
}
