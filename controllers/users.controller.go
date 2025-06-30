package controllers

import (
	"net/http"

	"github.com/fulsep/go-backend/models"
	"github.com/fulsep/go-backend/utils"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(ctx *gin.Context) {
	users, err := models.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get user from database",
		})
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List all users",
		Results: users,
	})
}

func DetailUser(ctx *gin.Context) {
	id, err := utils.GetPathInt(ctx, "id")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get id from path",
		})
	}

	user, err := models.FindOneUser(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get user from database",
		})
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Detail user",
		Results: user,
	})
}
