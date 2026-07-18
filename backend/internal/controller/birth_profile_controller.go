package controller

import (
	"destiny-backend/internal/dto"
	"destiny-backend/internal/model"
	"destiny-backend/internal/repository"
	"destiny-backend/pkg/response"
	"destiny-backend/pkg/utils"

	"github.com/labstack/echo/v4"
)

type BirthProfileController struct {
	repo *repository.BirthProfileRepo
}

func NewBirthProfileController(repo *repository.BirthProfileRepo) *BirthProfileController {
	return &BirthProfileController{repo: repo}
}

// @Summary List birth profiles
// @Description Get all birth profiles for the authenticated user
// @Tags BirthProfiles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profiles [get]
func (c *BirthProfileController) List(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	profiles, err := c.repo.ListByUserID(ctx.Request().Context(), claims.UserID)
	if err != nil {
		return response.Error(ctx, 500, "failed to fetch profiles")
	}
	return response.Success(ctx, profiles)
}

// @Summary Get birth profile by ID
// @Description Get a specific birth profile by its ID
// @Tags BirthProfiles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Birth Profile ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profiles/{id} [get]
func (c *BirthProfileController) GetByID(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	id := ctx.Param("id")

	profile, err := c.repo.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return response.Error(ctx, 500, "database error")
	}
	if profile == nil || profile.UserID != claims.UserID {
		return response.Error(ctx, 404, "birth profile not found")
	}
	return response.Success(ctx, profile)
}

// @Summary Create birth profile
// @Description Create a new birth profile for the authenticated user
// @Tags BirthProfiles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateBirthProfileRequest true "Create birth profile request"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profiles [post]
func (c *BirthProfileController) Create(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	var req dto.CreateBirthProfileRequest
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, 400, "invalid request body")
	}

	if req.FullName == "" || req.BirthDate == "" || req.Timezone == "" {
		return response.Error(ctx, 400, "full_name, birth_date and timezone are required")
	}

	bp := &model.BirthProfile{
		UserID:        claims.UserID,
		FullName:      req.FullName,
		Gender:        req.Gender,
		BirthDate:     req.BirthDate,
		BirthTime:     req.BirthTime,
		Timezone:      req.Timezone,
		Latitude:      req.Latitude,
		Longitude:     req.Longitude,
		City:          req.City,
		Country:       req.Country,
		BirthPlace:    req.BirthPlace,
		IsUnknownTime: req.IsUnknownTime,
		Note:          req.Note,
	}

	if err := c.repo.Create(ctx.Request().Context(), bp); err != nil {
		return response.Error(ctx, 500, "failed to create profile")
	}
	return response.Created(ctx, bp)
}

// @Summary Update birth profile
// @Description Update an existing birth profile
// @Tags BirthProfiles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Birth Profile ID"
// @Param request body dto.UpdateBirthProfileRequest true "Update birth profile request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profiles/{id} [put]
func (c *BirthProfileController) Update(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	id := ctx.Param("id")

	var req dto.UpdateBirthProfileRequest
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, 400, "invalid request body")
	}

	profile, err := c.repo.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return response.Error(ctx, 500, "database error")
	}
	if profile == nil || profile.UserID != claims.UserID {
		return response.Error(ctx, 404, "birth profile not found")
	}

	if req.FullName != nil {
		profile.FullName = *req.FullName
	}
	if req.Gender != nil {
		profile.Gender = req.Gender
	}
	if req.BirthDate != nil {
		profile.BirthDate = *req.BirthDate
	}
	if req.BirthTime != nil {
		profile.BirthTime = req.BirthTime
	}
	if req.Timezone != nil {
		profile.Timezone = *req.Timezone
	}
	if req.Latitude != nil {
		profile.Latitude = req.Latitude
	}
	if req.Longitude != nil {
		profile.Longitude = req.Longitude
	}
	if req.City != nil {
		profile.City = req.City
	}
	if req.Country != nil {
		profile.Country = req.Country
	}
	if req.BirthPlace != nil {
		profile.BirthPlace = req.BirthPlace
	}
	if req.IsUnknownTime != nil {
		profile.IsUnknownTime = *req.IsUnknownTime
	}
	if req.Note != nil {
		profile.Note = req.Note
	}

	if err := c.repo.Update(ctx.Request().Context(), profile); err != nil {
		return response.Error(ctx, 500, "failed to update profile")
	}
	return response.Success(ctx, profile)
}

// @Summary Delete birth profile
// @Description Delete a birth profile by ID
// @Tags BirthProfiles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Birth Profile ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profiles/{id} [delete]
func (c *BirthProfileController) Delete(ctx echo.Context) error {
	claims := ctx.Get("user").(*utils.Claims)
	id := ctx.Param("id")

	if err := c.repo.Delete(ctx.Request().Context(), id, claims.UserID); err != nil {
		return response.Error(ctx, 500, "failed to delete profile")
	}
	return response.Message(ctx, 200, "profile deleted successfully")
}
