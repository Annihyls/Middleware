package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllCollections() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Content)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetCollectionById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.User
	err = row.Scan(&collection.Id, &collection.Content)
	if err != nil {
		return nil, err
	}
	return &collection, err
}
