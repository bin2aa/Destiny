package controller

import (
	"destiny-backend/internal/repository"
	"destiny-backend/pkg/response"

	"github.com/labstack/echo/v4"
)

type LookupController struct {
	repo *repository.LookupRepo
}

func NewLookupController(repo *repository.LookupRepo) *LookupController {
	return &LookupController{repo: repo}
}

// @Summary Get zodiac signs
// @Description Get list of zodiac signs
// @Tags Lookup
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lookup/zodiac-signs [get]
func (c *LookupController) GetZodiacSigns(ctx echo.Context) error {
	data, err := c.repo.GetZodiacSigns(ctx.Request().Context())
	if err != nil {
		return response.Error(ctx, 500, "failed to fetch zodiac signs")
	}
	return response.Success(ctx, data)
}

// @Summary Get planets
// @Description Get list of planets
// @Tags Lookup
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lookup/planets [get]
func (c *LookupController) GetPlanets(ctx echo.Context) error {
	data, err := c.repo.GetPlanets(ctx.Request().Context())
	if err != nil {
		return response.Error(ctx, 500, "failed to fetch planets")
	}
	return response.Success(ctx, data)
}

// @Summary Get houses
// @Description Get list of houses
// @Tags Lookup
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lookup/houses [get]
func (c *LookupController) GetHouses(ctx echo.Context) error {
	data, err := c.repo.GetHouses(ctx.Request().Context())
	if err != nil {
		return response.Error(ctx, 500, "failed to fetch houses")
	}
	return response.Success(ctx, data)
}

// @Summary Get aspects
// @Description Get list of aspects
// @Tags Lookup
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lookup/aspects [get]
func (c *LookupController) GetAspects(ctx echo.Context) error {
	data, err := c.repo.GetAspects(ctx.Request().Context())
	if err != nil {
		return response.Error(ctx, 500, "failed to fetch aspects")
	}
	return response.Success(ctx, data)
}

// @Summary Get chinese zodiac
// @Description Get list of chinese zodiac signs
// @Tags Lookup
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lookup/chinese-zodiac [get]
func (c *LookupController) GetChineseZodiac(ctx echo.Context) error {
	data, err := c.repo.GetChineseZodiac(ctx.Request().Context())
	if err != nil {
		return response.Error(ctx, 500, "failed to fetch chinese zodiac")
	}
	return response.Success(ctx, data)
}

// @Summary Get five elements
// @Description Get list of five elements
// @Tags Lookup
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /lookup/five-elements [get]
func (c *LookupController) GetFiveElements(ctx echo.Context) error {
	data, err := c.repo.GetFiveElements(ctx.Request().Context())
	if err != nil {
		return response.Error(ctx, 500, "failed to fetch five elements")
	}
	return response.Success(ctx, data)
}
