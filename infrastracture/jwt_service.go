package infrastracture

import (
	"fmt"
	"librarybackend/config"
	"librarybackend/domain"

	"github.com/golang-jwt/jwt"
)

// TokenGenerator implementation
type TokenGeneratorImpl struct{}

// NewTokenGenerator creates a new TokenGenerator instance
func NewTokenGenerator() domain.TokenGenerator {
	return &TokenGeneratorImpl{}
}

// GenerateToken generates an access token for the user
func (tg *TokenGeneratorImpl) GenerateToken(user domain.User) (string, error) {
	env, err := config.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	accessTokenSecret := []byte(env.AccessTokenSecret)
	// accessTokenExpiryHour := env.AccessTokenExpiryHour

	claims := domain.JwtCustomClaims{
		UserID: user.ID.Hex(),
		// Role:   user.Role,
		Name: user.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(accessTokenSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

// GenerateRefreshToken generates a refresh token for the user
func (tg *TokenGeneratorImpl) GenerateRefreshToken(user domain.User) (string, error) {
	env, err := config.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	refreshTokenSecret := []byte(env.AccessTokenSecret)
	// refreshTokenExpiryHour := env.RefreshTokenExpiryHour

	claims := domain.JwtCustomClaims{
		UserID: user.ID.Hex(),
		// Role:   user.Role,
		Name: user.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(refreshTokenSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

// RefreshToken parses and verifies a refresh token and returns the user ID
func (tg *TokenGeneratorImpl) RefreshToken(tokenString string) (string, error) {
	env, err := config.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	refreshTokenSecret := []byte(env.AccessTokenSecret)

	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", err
	}

	return userID, nil
}
