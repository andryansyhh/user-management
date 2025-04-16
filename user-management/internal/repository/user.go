package repository

import (
	"context"
	"database/sql"
	"user-management/internal/domain/model"
)

type UserRepository interface {
	Fetch(ctx context.Context) ([]*model.User, error)
	Store(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Fetch(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT u.id, u.name, u.email, u.role_id, u.last_access, r.name
		FROM users u JOIN roles r ON u.role_id = r.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.RoleID, &u.LastAccess, &u.RoleName); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *userRepository) Store(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO users(name, email, password, role_id) VALUES($1, $2, $3, $4)`,
		user.Name, user.Email, user.Password, user.RoleID)
	return err
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4`,
		user.Name, user.Email, user.Password, user.ID)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}
