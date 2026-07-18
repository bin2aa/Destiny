package service

import (
	"context"
	"time"

	"destiny-backend/internal/dto"
	"destiny-backend/internal/model"
	"destiny-backend/internal/repository"
	"destiny-backend/pkg/errors"
	"destiny-backend/pkg/utils"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo      *repository.UserRepo
	redisClient   *redis.Client
	jwtSecret     string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewAuthService(userRepo *repository.UserRepo, redisClient *redis.Client, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:      userRepo,
		redisClient:   redisClient,
		jwtSecret:     jwtSecret,
		accessExpiry:  24 * time.Hour,
		refreshExpiry: 7 * 24 * time.Hour,
	}
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.TokenResponse, error) {
	// Check if email already exists
	existing, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NewInternalError("database error")
	}
	if existing != nil {
		return nil, errors.NewConflictError("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalError("failed to hash password")
	}

	user := &model.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Language:     "vi",
		Timezone:     "Asia/Ho_Chi_Minh",
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, errors.NewInternalError("failed to create user")
	}

	return s.generateTokens(user.ID, user.Email)
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.TokenResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NewInternalError("database error")
	}
	if user == nil {
		return nil, errors.NewUnauthorizedError("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.NewUnauthorizedError("invalid email or password")
	}

	return s.generateTokens(user.ID, user.Email)
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*dto.TokenResponse, error) {
	claims, err := utils.ValidateToken(refreshToken, s.jwtSecret)
	if err != nil {
		return nil, errors.NewUnauthorizedError("invalid or expired refresh token")
	}

	// Check if token is blacklisted in Redis
	if s.redisClient != nil {
		val, _ := s.redisClient.Exists(ctx, "blacklist:"+refreshToken).Result()
		if val > 0 {
			return nil, errors.NewUnauthorizedError("token has been revoked")
		}
	}

	return s.generateTokens(claims.UserID, claims.Email)
}

func (s *AuthService) Logout(ctx context.Context, accessToken, refreshToken string) error {
	if s.redisClient == nil {
		return nil
	}

	// Blacklist access token until expiry
	accessClaims, err := utils.ValidateToken(accessToken, s.jwtSecret)
	if err == nil {
		ttl := time.Until(accessClaims.ExpiresAt.Time)
		if ttl > 0 {
			s.redisClient.Set(ctx, "blacklist:"+accessToken, "1", ttl)
		}
	}

	// Blacklist refresh token
	refreshClaims, err := utils.ValidateToken(refreshToken, s.jwtSecret)
	if err == nil {
		ttl := time.Until(refreshClaims.ExpiresAt.Time)
		if ttl > 0 {
			s.redisClient.Set(ctx, "blacklist:"+refreshToken, "1", ttl)
		}
	}

	return nil
}

func (s *AuthService) generateTokens(userID, email string) (*dto.TokenResponse, error) {
	accessToken, err := utils.GenerateToken(userID, email, s.jwtSecret, s.accessExpiry)
	if err != nil {
		return nil, errors.NewInternalError("failed to generate access token")
	}

	refreshToken, err := utils.GenerateToken(userID, email, s.jwtSecret, s.refreshExpiry)
	if err != nil {
		return nil, errors.NewInternalError("failed to generate refresh token")
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(s.accessExpiry.Seconds()),
	}, nil
}
