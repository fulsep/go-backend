package routers

import "github.com/gin-gonic/gin"

func CombineRouters(r *gin.RouterGroup) {
	userRouter(r.Group("/users"))
}
