package dbrepo

import (
	"backend/internal/entities"
	"backend/internal/lib/e"
	"context"
	"database/sql"
	"errors"
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

func (db *PostgresDBRepo) Create(ctx context.Context, user entities.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT INTO users (first_name, last_name, email, password, created_at, updated_at, is_deleted) 
				 VALUES ($1, $2, $3, $4, now(), now() , FALSE)
				 RETURNING id`

	var userId int
	err := db.conn.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(&userId)

	if err != nil {
		return 0, e.WrapIfErr("failed to execute query", err)
	}

	return userId, nil
}

func (db *PostgresDBRepo) GetByID(ctx context.Context, id int) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at, is_deleted 
				 FROM users WHERE id = $1`

	var user entities.User
	err := db.conn.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user with not found")
	} else if err != nil {
		return nil, e.WrapIfErr("failed to execute query", err)
	}

	return &user, nil
}

func (db *PostgresDBRepo) Update(ctx context.Context, user entities.User) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `UPDATE users 
				 SET first_name = $1, last_name = $2, email = $3, password = $4, updated_at = now()
				 WHERE id = $5`

	if _, err := db.conn.ExecContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Id,
	); err != nil {
		return e.WrapIfErr("error executing query", err)
	}

	return nil
}

func (db *PostgresDBRepo) Delete(ctx context.Context, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `UPDATE users 
				 SET is_deleted = TRUE, updated_at = now()
				 WHERE id = $1`

	if _, err := db.conn.ExecContext(ctx, query, userId); err != nil {
		return e.WrapIfErr("error executing query", err)
	}

	return nil
}

func (db *PostgresDBRepo) List(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at, is_deleted 
				 FROM users LIMIT $1 OFFSET $2`

	var users []*entities.User
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

		users = append(users, &user)
	}

	return users, nil
}
