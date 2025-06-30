package controllers

import (
	"net/http"

	"github.com/fulsep/go-backend/utils"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List all users",
	})
}
