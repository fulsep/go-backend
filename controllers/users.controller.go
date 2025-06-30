package controllers

import (
	"net/http"

	"github.com/fulsep/go-backend/models"
	"github.com/fulsep/go-backend/utils"
	"github.com/gin-gonic/gin"
)

// @Description list all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Router /users [get]
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

// @Description list all users
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "string"
// @Router /users/{id} [get]
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
