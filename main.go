package main

import (
	"net/http"

	"github.com/fulsep/go-backend/routers"
	"github.com/fulsep/go-backend/utils"
	"github.com/gin-gonic/gin"
)

// @title           CRUD Users
// @version         1.0
// @description     This is a sample server celler server.
// @BasePath /

// @SecurityDefinitions.ApiKey Token
// @in header
// @name Authorization

func main() {
	app := gin.New()

	router := app.Group("/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, utils.Response{
			Success: true,
			Message: "Backend is running well",
		})
	})

	routers.CombineRouters(router)
	app.Run(":8888")
}
