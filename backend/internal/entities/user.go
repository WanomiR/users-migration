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

type ListConditions struct {
	NUsers int `json:"n_users,int" example:"5" binding:"required"`
	Limit  int `json:"limit,int" example:"10" binding:"required"`
	Offset int `json:"offset,int" example:"0" binding:"required"`
}
