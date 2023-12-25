package models
import (
    "github.com/gofrs/uuid"
)
type Rating struct {
    Id              *uuid.UUID      `json:"id"`
	Note            int             `json:"note"`
	Description     *string         `json:"description"`
	UserId          *uuid.UUID      `json:"id_user"`
    SongId          *uuid.UUID      `json:"id_song"`
}
