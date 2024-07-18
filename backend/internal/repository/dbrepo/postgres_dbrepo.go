package dbrepo

import (
	"backend/internal/entities"
	"backend/internal/lib/e"
	"context"
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	conn    *sql.DB
	timeout time.Duration
}

func WithDBTimeout(timeout time.Duration) func(*PostgresDBRepo) {
	return func(db *PostgresDBRepo) {
		db.timeout = timeout
	}

}

func NewPostgresDBRepo(conn *sql.DB, options ...func(repo *PostgresDBRepo)) *PostgresDBRepo {
	db := &PostgresDBRepo{
		conn:    conn,
		timeout: time.Second * 3, // default timout
	}

	for _, option := range options {
		option(db)
	}

	return db
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

func (db *PostgresDBRepo) List(ctx context.Context, limit, offset int) ([]entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at, is_deleted 
				 FROM users LIMIT $1 OFFSET $2`

	var users []entities.User
	rows, err := db.conn.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, e.WrapIfErr("error executing query", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsDeleted,
		)
		if err != nil {
			return nil, e.WrapIfErr("error scanning row", err)
		}

		users = append(users, user)
	}

	return users, nil
}
