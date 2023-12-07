package ratings

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"database/sql"
)

func GetAllRatings() ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.Note, &data.Description)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return ratings, err
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM ratings WHERE id=?", id.String())
	helpers.CloseDB(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.Note, &rating.Description)
	if err != nil {
		return nil, err
	}
	return &rating, err
}

func UpdateRating(id uuid.UUID, note int, description *string) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	if &description != nil {
	    _, err = db.Exec("UPDATE ratings SET note = ?, Description = ? WHERE id=?", note, &description, id.String())
	} else {
	    _, err = db.Exec("UPDATE ratings SET note = ?, Description = ? WHERE id=?", note, sql.NullString{}, id.String())
	}

	helpers.CloseDB(db)
	if err != nil {
    	return err
    }
	return nil
}

func DeleteRating(id uuid.UUID) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	db.Exec("DELETE FROM ratings WHERE id=?", id.String())
	helpers.CloseDB(db)
	if err != nil {
    	return err
    }
	return nil
}

func CreateRating(id uuid.UUID, note int, description *string) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	if &description != nil {
	    _, err = db.Exec("INSERT INTO ratings VALUES (?, ?, ?)", id.String(), note, &description)
	} else {
	    _, err = db.Exec("INSERT INTO ratings (id, note) VALUES (?, ?)", id.String(), note)
	}

	helpers.CloseDB(db)
	if err != nil {
    	return err
    }
	return nil
}