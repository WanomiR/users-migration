package dbrepo

import (
	"backend/internal/entities"
	"context"
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	conn *sql.DB
}

func NewPostgresDBRepo(conn *sql.DB) *PostgresDBRepo {
	return &PostgresDBRepo{conn: conn}
}

func (db *PostgresDBRepo) Connection() *sql.DB {
	return db.conn
}

func (db *PostgresDBRepo) Create(ctx context.Context, user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (db *PostgresDBRepo) GetByID(ctx context.Context, id int) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (db *PostgresDBRepo) Update(ctx context.Context, user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (db *PostgresDBRepo) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (db *PostgresDBRepo) List(ctx context.Context, c entities.ListConditions) ([]entities.User, error) {
	//TODO implement me
	panic("implement me")
}
