package repository

import (
	"context"
	"errors"
	"pathshala/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository struct {
	DB *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) GetAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	query := `SELECT id, name, email, password, role, created_at FROM admins WHERE email = $1`

	err := r.DB.QueryRow(context.Background(), query, email).Scan(
		&admin.ID, &admin.Name, &admin.Email, &admin.Password, &admin.Role, &admin.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("admin not found")
		}
		return nil, err
	}
	return &admin, nil
}
