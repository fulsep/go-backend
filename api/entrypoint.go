package main

import (
	"net/http"

	"github.com/fulsep/go-backend/routers"
	"github.com/fulsep/go-backend/utils"
	"github.com/gin-gonic/gin"
)

var App *gin.Engine

func main() {
	App = gin.New()

	router := App.Group("/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, utils.Response{
			Success: true,
			Message: "Backend is running well",
		})
	})

	routers.CombineRouters(router)
	App.Run(":8888")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
