package repositories

import (
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

func (repository *userRepositoryImpl) Insert(usr *entity.User) (*entity.User, error) {
	query := `
	INSERT INTO users 
	(id, username, password, contact, role, created_at, updated_at) 
	VALUES 
	(?, ?, ?, ?, ?, ?, ?)
	`

	_, err := repository.database.Exec(
		query,
		usr.ID,
		usr.Username,
		usr.Password,
		usr.Contact,
		usr.Role,
		usr.CreatedAt,
		usr.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (repository *userRepositoryImpl) FindAll() ([]*entity.User, error) {
	query := "SELECT * FROM users"

	rows, err := repository.database.Query(query)
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
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfUsers = append(listOfUsers, &user)
	}
	// rerr := rows.Close()
	// if rerr != nil {
	// 	return nil, err
	// }

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfUsers, nil
}

func (repository *userRepositoryImpl) FindByID(usrID uuid.UUID) (*entity.User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	row := repository.database.QueryRow(query, usrID)

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Contact,
		&user.Role,
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

func (repository *userRepositoryImpl) GetByUsername(name string) ([]*entity.User, error) {
	query := "SELECT * FROM users WHERE username LIKE ?"
	name = name + "%"

	rows, err := repository.database.Query(query, name)
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
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfUsers = append(listOfUsers, &user)
	}
	// rerr := rows.Close()
	// if rerr != nil {
	// 	return nil, err
	// }

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfUsers, nil
}

func (repository *userRepositoryImpl) FindByUsername(name string) (*entity.User, error) {
	query := "SELECT * FROM users WHERE username = ?"

	rows := repository.database.QueryRow(query, name)

	var user entity.User
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Contact,
		&user.Role,
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

func (repository *userRepositoryImpl) Update(usr *entity.User) error {
	query := `
	UPDATE users 
	SET username = ?, password = ?, contact = ?, role = ?, updated_at = ? 
	WHERE id = ?
	`

	_, err := repository.database.Exec(
		query,
		usr.Username,
		usr.Password,
		usr.Contact,
		usr.Role,
		usr.UpdatedAt,
		usr.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}

func (repository *userRepositoryImpl) Delete(usrID uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := repository.database.Exec(query, usrID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}
