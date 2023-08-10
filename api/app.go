package api

import "github.com/gin-gonic/gin"

func HandleRoutes(addr string) {
	app := gin.New()

	app.GET("/:query", handleGetPron)

	app.Run(addr)
}
