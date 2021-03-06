package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mtdx/case-api/routes"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := routes.SetupRouter()
	router.Run()
}
