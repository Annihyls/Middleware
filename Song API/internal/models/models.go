package models

import (
	"github.com/gofrs/uuid"
)

/*type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Content string     `json:"content"`
}*/

type Song struct {
	ID     *uuid.UUID `json:"id"`
	Title  string     `json:"title"`
	Artist string     `json:"artist"`
}
