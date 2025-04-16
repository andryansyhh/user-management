package repository

import (
	"context"
	"database/sql"
	"errors"

	"user-management/internal/domain/model"
)

type AuthRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT u.id, u.name, u.email, u.password, u.role_id, r.name 
			  FROM users u
			  JOIN roles r ON u.role_id = r.id
			  WHERE u.email = $1`

	row := r.db.QueryRowContext(ctx, query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.RoleName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
