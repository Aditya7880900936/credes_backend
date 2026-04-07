package models

import "time"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID        int64     `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password"` // hashed
	FullName  string    `db:"full_name" json:"full_name"`
	Role      Role      `db:"role" json:"role"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}