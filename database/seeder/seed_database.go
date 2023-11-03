package seeder

import (
	"database/sql"
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/google/uuid"
)

func RolesSeed(database *sql.DB) error {
	query1 := "SELECT COUNT(*) FROM roles"
	var count int

	err := database.QueryRow(query1).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		query2 := `
			INSERT INTO roles
			(id, name, created_at, updated_at)
			VALUES
			(?, ?, ?, ?)
		`

		listOfRoles := []entity.Role{
			{
				ID:        1,
				Name:      "Warehouse Admin",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        2,
				Name:      "Warehouse Staff",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        3,
				Name:      "External",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		for _, role := range listOfRoles {
			_, err := database.Exec(
				query2,
				role.ID,
				role.Name,
				role.CreatedAt,
				role.UpdatedAt,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func AdminUserSeed(database *sql.DB) error {
	query1 := "SELECT COUNT(*) FROM users"
	var count int

	err := database.QueryRow(query1).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		hasedPassword, err := helpers.HashPassword("admin123")
		if err != nil {
			return err
		}

		admin := entity.User{
			ID:        uuid.New(),
			Username:  "admin",
			Password:  hasedPassword,
			Contact:   "admin",
			RoleID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Role:      nil,
		}

		query2 := `
			INSERT INTO users
			(id, username, password, contact, role_id, created_at, updated_at)
			VALUES
			(?, ?, ?, ?, ?, ?, ?)
		`

		_, err = database.Exec(
			query2,
			admin.ID,
			admin.Username,
			admin.Password,
			admin.Contact,
			admin.RoleID,
			admin.CreatedAt,
			admin.UpdatedAt,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
