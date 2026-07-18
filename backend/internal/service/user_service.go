package service

import (
	"context"
	"time"

	"destiny-backend/internal/dto"
	"destiny-backend/internal/model"
	"destiny-backend/internal/repository"
	"destiny-backend/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetProfile(ctx context.Context, userID string) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.NewInternalError("database error")
	}
	if user == nil {
		return nil, errors.NewNotFoundError("user")
	}

	return toUserResponse(user), nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userID string, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.NewInternalError("database error")
	}
	if user == nil {
		return nil, errors.NewNotFoundError("user")
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Avatar != nil {
		user.Avatar = req.Avatar
	}
	if req.Language != nil {
		user.Language = *req.Language
	}
	if req.Timezone != nil {
		user.Timezone = *req.Timezone
	}
	if req.Country != nil {
		user.Country = req.Country
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, errors.NewInternalError("failed to update user")
	}

	return toUserResponse(user), nil
}

func (s *UserService) ChangePassword(ctx context.Context, userID string, req *dto.ChangePasswordRequest) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.NewInternalError("database error")
	}
	if user == nil {
		return errors.NewNotFoundError("user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword)); err != nil {
		return errors.NewBadRequestError("current password is incorrect")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.NewInternalError("failed to hash password")
	}

	return s.userRepo.UpdatePassword(ctx, userID, string(hashedPassword))
}

func toUserResponse(user *model.User) *dto.UserResponse {
	resp := &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Language:  user.Language,
		Timezone:  user.Timezone,
		Country:   user.Country,
		IsPremium: false,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	if user.PremiumUntil != nil && user.PremiumUntil.After(time.Now()) {
		resp.IsPremium = true
		until := user.PremiumUntil.Format(time.RFC3339)
		resp.PremiumUntil = &until
	}

	return resp
}
