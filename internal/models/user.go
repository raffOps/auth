package models

import "time"

type LoginHistory struct {
	LoginTime  time.Time `json:"login_time"`
	LogoutTime time.Time `json:"logout_time"`
}

type User struct {
	Id           string         `json:"id"`
	Username     string         `json:"name" validate:"required,min=5,max=100"`
	Email        string         `json:"email"`
	AuthType     string         `json:"auth_type" validate:"required,oneof=GOOGLE GITHUB"`
	UserGroup    string         `json:"user_group" validate:"required,oneof=ADMIN USER"`
	Status       string         `json:"status" validate:"required,oneof=ACTIVE INACTIVE"`
	LoginHistory []LoginHistory `json:"login_histories"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    time.Time      `json:"deleted_at"`
}
