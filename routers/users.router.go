package routers

import (
	"github.com/fulsep/go-backend/controllers"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllUsers)
	r.GET("/:id", controllers.DetailUser)
}
