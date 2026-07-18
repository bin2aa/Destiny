package routes

import (
	"time"

	"destiny-backend/internal/controller"
	"destiny-backend/internal/middleware"
	"destiny-backend/internal/model"
	"destiny-backend/internal/repository"
	"destiny-backend/internal/service"
	"destiny-backend/pkg/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func Setup(e *echo.Echo, pool *pgxpool.Pool, redisClient *redis.Client, jwtSecret string) {
	// Repositories
	userRepo := repository.NewUserRepo(pool)
	birthProfileRepo := repository.NewBirthProfileRepo(pool)
	lookupRepo := repository.NewLookupRepo(pool)
	billingRepo := repository.NewBillingRepo(pool)
	chatRepo := repository.NewChatRepo(pool)
	horoscopeRepo := repository.NewHoroscopeRepo(pool)

	// Controllers - services are embedded in controllers
	lookupCtrl := controller.NewLookupController(lookupRepo)
	bpCtrl := controller.NewBirthProfileController(birthProfileRepo)

	// JWT middleware
	jwtMiddleware := middleware.JWTConfig(jwtSecret)

	// API v1
	v1 := e.Group("/api/v1")

	// Public routes
	v1.POST("/auth/register", func(c echo.Context) error {
		authCtrl := controller.NewAuthController(service.NewAuthService(userRepo, redisClient, jwtSecret))
		return authCtrl.Register(c)
	})
	v1.POST("/auth/login", func(c echo.Context) error {
		authCtrl := controller.NewAuthController(service.NewAuthService(userRepo, redisClient, jwtSecret))
		return authCtrl.Login(c)
	})
	v1.POST("/auth/refresh", func(c echo.Context) error {
		authCtrl := controller.NewAuthController(service.NewAuthService(userRepo, redisClient, jwtSecret))
		return authCtrl.RefreshToken(c)
	})

	// Public lookup routes
	v1.GET("/lookup/zodiac-signs", lookupCtrl.GetZodiacSigns)
	v1.GET("/lookup/planets", lookupCtrl.GetPlanets)
	v1.GET("/lookup/houses", lookupCtrl.GetHouses)
	v1.GET("/lookup/aspects", lookupCtrl.GetAspects)
	v1.GET("/lookup/chinese-zodiac", lookupCtrl.GetChineseZodiac)
	v1.GET("/lookup/five-elements", lookupCtrl.GetFiveElements)

	// Protected routes
	protected := v1.Group("")
	protected.Use(jwtMiddleware)

	// Auth
	protected.POST("/auth/logout", func(c echo.Context) error {
		authCtrl := controller.NewAuthController(service.NewAuthService(userRepo, redisClient, jwtSecret))
		return authCtrl.Logout(c)
	})

	// User
	userCtrl := controller.NewUserController(service.NewUserService(userRepo))
	protected.GET("/users/me", userCtrl.GetProfile)
	protected.PUT("/users/me", userCtrl.UpdateProfile)
	protected.PUT("/users/me/password", userCtrl.ChangePassword)

	// Birth Profiles
	protected.GET("/profiles", bpCtrl.List)
	protected.POST("/profiles", bpCtrl.Create)
	protected.GET("/profiles/:id", bpCtrl.GetByID)
	protected.PUT("/profiles/:id", bpCtrl.Update)
	protected.DELETE("/profiles/:id", bpCtrl.Delete)

	// Billing
	protected.GET("/plans", func(c echo.Context) error {
		plans, err := billingRepo.GetPlans(c.Request().Context())
		if err != nil {
			return c.JSON(500, map[string]interface{}{"success": false, "message": "failed to fetch plans"})
		}
		return c.JSON(200, map[string]interface{}{"success": true, "data": plans})
	})
	protected.GET("/subscriptions", func(c echo.Context) error {
		claims := c.Get("user").(*utils.Claims)
		subs, err := billingRepo.GetSubscriptionsByUserID(c.Request().Context(), claims.UserID)
		if err != nil {
			return c.JSON(500, map[string]interface{}{"success": false, "message": "failed to fetch subscriptions"})
		}
		return c.JSON(200, map[string]interface{}{"success": true, "data": subs})
	})

	// Chat
	protected.GET("/chat", func(c echo.Context) error {
		claims := c.Get("user").(*utils.Claims)
		chats, err := chatRepo.ListByUserID(c.Request().Context(), claims.UserID)
		if err != nil {
			return c.JSON(500, map[string]interface{}{"success": false, "message": "failed to fetch chats"})
		}
		return c.JSON(200, map[string]interface{}{"success": true, "data": chats})
	})
	protected.POST("/chat", func(c echo.Context) error {
		claims := c.Get("user").(*utils.Claims)
		var req struct {
			Title          *string `json:"title"`
			BirthProfileID *string `json:"birth_profile_id"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]interface{}{"success": false, "message": "invalid request"})
		}
		chat := &model.AIChat{
			UserID:         claims.UserID,
			BirthProfileID: req.BirthProfileID,
			Title:          req.Title,
		}
		if err := chatRepo.CreateChat(c.Request().Context(), chat); err != nil {
			return c.JSON(500, map[string]interface{}{"success": false, "message": "failed to create chat"})
		}
		return c.JSON(201, map[string]interface{}{"success": true, "data": chat})
	})
	protected.GET("/chat/:id/messages", func(c echo.Context) error {
		msgs, err := chatRepo.GetMessagesByChatID(c.Request().Context(), c.Param("id"))
		if err != nil {
			return c.JSON(500, map[string]interface{}{"success": false, "message": "failed to fetch messages"})
		}
		return c.JSON(200, map[string]interface{}{"success": true, "data": msgs})
	})

	// Horoscope
	protected.GET("/horoscope/today", func(c echo.Context) error {
		claims := c.Get("user").(*utils.Claims)
		profiles, _ := birthProfileRepo.ListByUserID(c.Request().Context(), claims.UserID)
		if len(profiles) == 0 {
			return c.JSON(200, map[string]interface{}{"success": true, "data": nil})
		}
		today := time.Now().Format("2006-01-02")
		h, err := horoscopeRepo.GetByProfileAndDate(c.Request().Context(), profiles[0].ID, today)
		if err != nil {
			return c.JSON(500, map[string]interface{}{"success": false, "message": "failed to fetch horoscope"})
		}
		return c.JSON(200, map[string]interface{}{"success": true, "data": h})
	})
}
