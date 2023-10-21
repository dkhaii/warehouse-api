package entity

import "time"

type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      []User
}

type RoleFiltered struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
