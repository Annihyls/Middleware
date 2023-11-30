package models
import (
    "github.com/gofrs/uuid"
)
type Rating struct {
    //id de l'utilisateur
    Id          *uuid.UUID       `json:"id"`
	Note        int             `json:"note"`
	Description *string         `json:"description"`
}
