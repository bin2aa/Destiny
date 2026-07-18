package utils

import (
	"math"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims represents JWT claims.
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token.
func GenerateToken(userID, email, secret string, expiry time.Duration) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken parses and validates a JWT token string.
func ValidateToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// Pagination holds pagination parameters.
type Pagination struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}

// PaginationInfo calculates pagination metadata.
func PaginationInfo(page, perPage, total int) (offset int, totalPages int) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 20
	}
	if perPage > 100 {
		perPage = 100
	}

	offset = (page - 1) * perPage
	totalPages = int(math.Ceil(float64(total) / float64(perPage)))
	return offset, totalPages
}

// Ptr returns a pointer to a copy of the given value.
func Ptr[T any](v T) *T {
	return &v
}
