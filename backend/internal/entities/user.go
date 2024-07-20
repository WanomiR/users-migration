package entities

import "time"

type User struct {
	Id        int       `json:"id,int"`
	FirstName string    `json:"first_name" example:"John" binding:"required"`
	LastName  string    `json:"last_name" example:"Smith" binding:"required"`
	Email     string    `json:"email" example:"john.smith@gmail.com" binding:"required"`
	Password  string    `json:"password" example:"password" binding:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsDeleted bool      `json:"-"`
}
