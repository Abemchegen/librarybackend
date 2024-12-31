package domain

import (
	"context"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
	// Role     string             `bson:"role" json:"role"`
	// Image string `bson:"image" json:"image"`
}

type TokenGenerator interface {
	GenerateToken(user User) (string, error)
	GenerateRefreshToken(user User) (string, error)
	RefreshToken(token string) (string, error)
}
type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Name   string `json:"name"`

	jwt.StandardClaims
}

type UserUseCase interface {
	CreateAccount(user User) (User, ErrorResponse)
	Login(user User) (LoginSucessResponse, ErrorResponse)
	GetByID(id string) (User, ErrorResponse)

	UpdateProfile(id string, user User) (User, ErrorResponse)

	GetAllUser() ([]User, error)

	GetUserByID(id string) (User, error)
}

type UserRepository interface {
	CreateAccount(user User) (User, error)
	Login(user User) (User, error)
	GetAllUserByEmial(email string) (User, error)
	GetByID(id string) (User, error)
	UpdateProfile(id string, user User) (User, error)

	GetAllUser() ([]User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
}
