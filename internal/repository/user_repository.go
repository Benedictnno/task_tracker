package repository

import (
	"context"

	"todo-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow(ctx,
		`INSERT INTO users (email, password)
		 VALUES ($1, $2)
		 RETURNING id, email, password, created_at`,
		email,
		password,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow(
		ctx,
		`SELECT id, email, password, created_at
		 FROM users
		 WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}