package routes

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/handlers"
	"github.com/mtdx/case-api/middleware"
)

var router *gin.Engine

func initRoutes() {
	authMiddleware := middleware.Jwt()
	// limit simultaneous connections
	// router.Use(middleware.LimitMax(200))

	// Gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/", handlers.Index)
		apiv1.POST("/login", authMiddleware.LoginHandler) // TODO: implement
		restricted := apiv1.Group("/").Use(authMiddleware.MiddlewareFunc())
		{
			restricted.GET("/refresh-token", authMiddleware.RefreshHandler)
			restricted.GET("/authenticated", handlers.Authenticated)
		}

	}
}

// SetupRouter ...
func SetupRouter() *gin.Engine {
	router = gin.Default()

	initRoutes()

	return router
}
