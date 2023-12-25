package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id      *uuid.UUID `json:"id"`
	Content string     `json:"content"`
	Nom     string     `json:"nom"`
	Prenom  string     `json:"prenom"`
}
