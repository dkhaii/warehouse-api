package repository

import (
	"database/sql"

	"github.com/dkhaii/warehouse-api/domain/user"
	"github.com/google/uuid"
)

type userRepositoryImpl struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{
		database: database,
	}
}

func (repository *userRepositoryImpl) Insert(usr user.UserEntity) (user.UserEntity, error) {
	query := "INSERT INTO users (id, name, contact, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	_, err := repository.database.Exec(
		query,
		usr.ID,
		usr.Name,
		usr.Contact,
		usr.Role,
		usr.CreatedAt,
		usr.UpdatedAt,
	)
	if err != nil {
		return user.UserEntity{}, err
	}

	return usr, nil
}

func (repository *userRepositoryImpl) FindAll() ([]user.UserEntity, error) {
	query := "SELECT * FROM users"

	users, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer users.Close()

	var listOfUsers []user.UserEntity

	for users.Next() {
		var u user.UserEntity

		err := users.Scan(&u.ID, &u.Name, &u.Contact, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		listOfUsers = append(listOfUsers, u)
	}

	err = users.Err()
	if err != nil {
		return nil, err
	}

	return listOfUsers, nil
}

func (repository *userRepositoryImpl) FindByID(usrID uuid.UUID) (user.UserEntity, error) {
	query := "SELECT * FROM users WHERE id = ?"

	sqlResult := repository.database.QueryRow(query, usrID)

	var user user.UserEntity
	err := sqlResult.Scan(&user.ID, &user.Name, &user.Contact, &user.Contact, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindByName(usrName string) ([]user.UserEntity, error) {
	query := "SELECT * FROM users WHERE name = ?"

	users, err := repository.database.Query(query, usrName)
	if err != nil {
		return nil, err
	}
	defer users.Close()

	var listOfUsers []user.UserEntity

	for users.Next() {
		var u user.UserEntity

		err := users.Scan(&u.ID, &u.Name, &u.Contact, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		listOfUsers = append(listOfUsers, u)
	}

	err = users.Err()
	if err != nil {
		return nil, err
	}

	return listOfUsers, nil
}

func (repository *userRepositoryImpl) Update(usr user.UserEntity) error {
	query := "UPDATE users SET name = ?, contact = ?, role = ?, updated_at = ? WHERE id = ?"

	_, err := repository.database.Exec(query, usr.Name, usr.Contact, usr.Role, usr.UpdatedAt, usr.ID)
	if err != nil {
		return err
	}
	defer repository.database.Close()

	return nil
}

func (repository *userRepositoryImpl) Delete(usrID user.UserEntity) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := repository.database.Exec(query, usrID)
	if err != nil {
		return nil
	}

	defer repository.database.Close()

	return nil
}

