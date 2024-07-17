package repository

import (
	"backend/internal/entities"
	"context"
	"database/sql"
)

type UserRepository interface {
	Connection() *sql.DB
	Create(ctx context.Context, user entities.User) error // which is insert
	GetByID(ctx context.Context, id int) (entities.User, error)
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, c entities.ListConditions) ([]entities.User, error)
	// Другие методы, необходимые для работы с пользователями
}
