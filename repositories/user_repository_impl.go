package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type userRepositoryImpl struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepositoryImpl{
		database: database,
	}
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, usr *entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users 
		(id, username, password, contact, role_id, created_at, updated_at) 
		VALUES 
		(?, ?, ?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		usr.ID,
		usr.Username,
		usr.Password,
		usr.Contact,
		usr.RoleID,
		usr.CreatedAt,
		usr.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (repository *userRepositoryImpl) FindAll(ctx context.Context) ([]*entity.User, error) {
	query := "SELECT * FROM users"

	rows, err := repository.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listOfUsers []*entity.User

	for rows.Next() {
		var user entity.User

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Contact,
			&user.RoleID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfUsers = append(listOfUsers, &user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfUsers, nil
}

func (repository *userRepositoryImpl) FindByID(ctx context.Context, usrID uuid.UUID) (*entity.User, error) {
	var user entity.User

	query := "SELECT * FROM users WHERE id = ?"

	err := repository.database.QueryRowContext(ctx, query, usrID).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Contact,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (repository *userRepositoryImpl) FindCompleteByID(ctx context.Context, usrID uuid.UUID) (*entity.User, error) {
	var user entity.User
	var role entity.RoleFiltered

	query := `
		SELECT usr.id, usr.username, usr.contact, usr.role_id, usr.created_at, usr. updated_at,
		r.id, r.name 
		FROM users usr
		LEFT JOIN roles r
		ON r.id = usr.role_id
		WHERE usr.id = ?
	`

	err := repository.database.QueryRowContext(ctx, query, usrID).Scan(
		&user.ID,
		&user.Username,
		&user.Contact,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&role.ID,
		&role.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	user.Role = &role

	return &user, nil
}

func (repository *userRepositoryImpl) GetByUsername(ctx context.Context, name string) ([]*entity.User, error) {
	query := "SELECT * FROM users WHERE username LIKE ?"
	name = name + "%"

	rows, err := repository.database.QueryContext(ctx, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	var listOfUsers []*entity.User

	for rows.Next() {
		var user entity.User

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Contact,
			&user.RoleID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfUsers = append(listOfUsers, &user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfUsers, nil
}

func (repository *userRepositoryImpl) FindByUsername(ctx context.Context, name string) (*entity.User, error) {
	query := "SELECT * FROM users WHERE username = ?"

	rows := repository.database.QueryRowContext(ctx, query, name)

	var user entity.User
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Contact,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (repository *userRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, usr *entity.User) (*entity.User, error) {
	query := `
		UPDATE users 
		SET username = ?, password = ?, contact = ?, role_id = ?, updated_at = ? 
		WHERE id = ?
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		usr.Username,
		usr.Password,
		usr.Contact,
		usr.RoleID,
		usr.UpdatedAt,
		usr.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return usr, nil
}

func (repository *userRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, usrID uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, usrID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}
