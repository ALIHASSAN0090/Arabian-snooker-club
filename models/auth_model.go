package models

import "github.com/golang-jwt/jwt"

type ResponseWithoutData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
type SignUpReq struct {
	UserName    string `json:"user_name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Role        string `json:"role,omitempty"`
	Address     string `json:"address,omitempty" validate:"required"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtUser struct {
	UserID *int64  `json:"id"`
	Email  *string `json:"email"`
	Name   *string `json:"name"`
	Role   *string `json:"role"`
}

type Admin struct {
	Id       uint    `json:"id" db:"id"`
	Email    *string `json:"email" db:"email"`
	Password string  `json:"password" db:"password"`
	Is_admin *bool   `json:"is_admin" db:"is_admin"`
}

type CustomClaims struct {
	Id    uint    `json:"id"`
	Email *string `json:"email"`
	Name  *string `json:"name"`
	Role  *string `json:"role"`
	jwt.StandardClaims
}
