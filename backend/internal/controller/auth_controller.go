package controller

import (
	"destiny-backend/internal/dto"
	"destiny-backend/internal/service"
	"destiny-backend/pkg/errors"
	"destiny-backend/pkg/response"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// @Summary Register a new user
// @Description Create a new user account and return access/refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register request"
// @Success 201 {object} dto.TokenResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func (c *AuthController) Register(ctx echo.Context) error {
	var req dto.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, 400, "invalid request body")
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return response.Error(ctx, 400, "name, email and password are required")
	}

	tokens, err := c.authService.Register(ctx.Request().Context(), &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return response.Error(ctx, appErr.Code, appErr.Message)
		}
		return response.Error(ctx, 500, "internal server error")
	}

	return response.Created(ctx, tokens)
}

// @Summary Login user
// @Description Authenticate user and return access/refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/login [post]
func (c *AuthController) Login(ctx echo.Context) error {
	var req dto.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, 400, "invalid request body")
	}

	tokens, err := c.authService.Login(ctx.Request().Context(), &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return response.Error(ctx, appErr.Code, appErr.Message)
		}
		return response.Error(ctx, 500, "internal server error")
	}

	return response.Success(ctx, tokens)
}

// @Summary Refresh access token
// @Description Get a new access token using a refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/refresh [post]
func (c *AuthController) RefreshToken(ctx echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, 400, "invalid request body")
	}

	tokens, err := c.authService.RefreshToken(ctx.Request().Context(), req.RefreshToken)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return response.Error(ctx, appErr.Code, appErr.Message)
		}
		return response.Error(ctx, 500, "internal server error")
	}

	return response.Success(ctx, tokens)
}

// @Summary Logout user
// @Description Invalidate access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx echo.Context) error {
	accessToken := ctx.Request().Header.Get("Authorization")
	if len(accessToken) > 7 {
		accessToken = accessToken[7:] // Strip "Bearer "
	}

	var req dto.RefreshTokenRequest
	if err := ctx.Bind(&req); err == nil && req.RefreshToken != "" {
		c.authService.Logout(ctx.Request().Context(), accessToken, req.RefreshToken)
	}

	return response.Message(ctx, 200, "logged out successfully")
}
