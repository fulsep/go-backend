package api

import (
	"net/http"

	"github.com/fulsep/go-backend/routers"
	"github.com/fulsep/go-backend/utils"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	app := gin.New()

	router := app.Group("/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, utils.Response{
			Success: true,
			Message: "Backend is running well",
		})
	})

	routers.CombineRouters(router)

	app.ServeHTTP(w, r)
}
