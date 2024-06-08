package models

import "time"

type Comentario struct {
	Id        int       `json:"id" db:"id"`
	Timestamp time.Time `json:"time" db:"time"`
	Comment   string    `json:"comment" db:"comment"`
	Reactions uint      `json:"reactions" db:"reactions"`
}

type Usuario struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
