package users

import (
	"database/sql"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllUsers() ([]models.User, error) {
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
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Prenom, &data.Nom)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.Id, &user.Prenom, &user.Nom)
	if err != nil {
		return nil, err
	}
	return &user, err
}
func UpdateUser(id uuid.UUID, prenom string, nom string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	if &nom != nil {
		_, err = db.Exec("UPDATE users SET prenom = ?, nom = ? WHERE id=?", prenom, &nom, id.String())
	} else {
		_, err = db.Exec("UPDATE users SET prenom = ?, nom = ? WHERE id=?", nom, sql.NullString{}, id.String())
	}

	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	db.Exec("DELETE FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(id uuid.UUID, prenom string, nom string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	if &nom != nil {
		_, err = db.Exec("INSERT INTO users VALUES (?, ?, ?)", id.String(), prenom, &nom)
	} else {
		_, err = db.Exec("INSERT INTO users (id, note) VALUES (?, ?)", id.String(), prenom)
	}

	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return nil
}
