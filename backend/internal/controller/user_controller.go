package controller

import (
	"destiny-backend/internal/dto"
	"destiny-backend/internal/service"
	"destiny-backend/pkg/errors"
	"destiny-backend/pkg/response"
	"destiny-backend/pkg/utils"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// @Summary Get current user profile
// @Description Get the authenticated user's profile information
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/me [get]
func (c *UserController) GetProfile(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	user, err := c.userService.GetProfile(ctx.Request().Context(), claims.UserID)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return response.Error(ctx, appErr.Code, appErr.Message)
		}
		return response.Error(ctx, 500, "internal server error")
	}
	return response.Success(ctx, user)
}

// @Summary Update user profile
// @Description Update the authenticated user's profile information
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateUserRequest true "Update user request"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/me [put]
func (c *UserController) UpdateProfile(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	var req dto.UpdateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, 400, "invalid request body")
	}

	user, err := c.userService.UpdateProfile(ctx.Request().Context(), claims.UserID, &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return response.Error(ctx, appErr.Code, appErr.Message)
		}
		return response.Error(ctx, 500, "internal server error")
	}
	return response.Success(ctx, user)
}

// @Summary Change password
// @Description Change the authenticated user's password
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ChangePasswordRequest true "Change password request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/me/password [put]
func (c *UserController) ChangePassword(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	var req dto.ChangePasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, 400, "invalid request body")
	}

	if err := c.userService.ChangePassword(ctx.Request().Context(), claims.UserID, &req); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return response.Error(ctx, appErr.Code, appErr.Message)
		}
		return response.Error(ctx, 500, "internal server error")
	}
	return response.Message(ctx, 200, "password changed successfully")
}
