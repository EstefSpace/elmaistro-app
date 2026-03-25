package models

import "time"

type Customer struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
