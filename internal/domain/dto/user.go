package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserDTO struct {
	Name     string             `json:"name"`
	Email    string             `json:"email" required:"true"`
	Password string             `json:"password" required:"true"`
	CartID   primitive.ObjectID `json:"cartID"`
}

type UpdateUserDTO struct {
	Name   string              `json:"name"`
	CartID *primitive.ObjectID `json:"cartID"`
}

type UpdateUserInput struct {
	Name   string              `json:"name"`
	CartID *primitive.ObjectID `json:"cartID"`
}

type SignUpDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"  binding:"required"`
}
