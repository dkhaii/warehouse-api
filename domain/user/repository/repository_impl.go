package repository

import (
	"database/sql"

	"github.com/dkhaii/warehouse-api/domain/user"
)

type userRepositoryImpl struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{
		database: database,
	}
}

func (repository *userRepositoryImpl) Insert(user user.UserEntity) (user.UserEntity, error) {
	query := "INSERT INTO users (id, name, contact, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	
	_, err := repository.database.Exec(
		query,
		user.ID,
		user.Name,
		user.Contact,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}


func (repository *userRepositoryImpl) FindAll() ([]user.UserEntity, error) {
	query := "SELECT * FROM users"

	users, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer users.Close()

	var listOfUsers []user.UserEntity

	for users.Next(){
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