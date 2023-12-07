package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllSongss() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.ID, &data.Title, &data.Artist)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.ID, &song.Title, &song.Artist)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func PostSong(title, artist string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Generate a new UUID for the new song
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Create a new song object with the provided data and the generated ID
	newSong := models.Song{
		ID:     &newID,
		Title:  title,
		Artist: artist,
	}

	// Execute the INSERT query to add the new song to the database
	_, err = db.Exec("INSERT INTO songs (id, title, artist) VALUES (?, ?, ?)",
		newSong.ID.String(), newSong.Title, newSong.Artist)

	if err != nil {
		return nil, err
	}

	return &newSong, nil
}

func UpdateSong(id uuid.UUID, newTitle, newArtist string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	// Execute the UPDATE query to modify the existing song in the SQLite database
	_, err = db.Exec("UPDATE songs SET title=?, artist=? WHERE id=?",
		newTitle, newArtist, id.String())

	return err
}

func DeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	// Execute the DELETE query to modify the existing song in the SQLite database
	_, err = db.Exec("DELETE FROM songs WHERE id=?",
		id.String())

	return err
}
