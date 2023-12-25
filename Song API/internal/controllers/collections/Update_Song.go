package collections

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
)

// UpdateSongHandler handles HTTP requests to update an existing song in the collection.
// @Tags         collections
// @Summary      Update an existing song in the collection.
// @Description  Update an existing song in the collection.
// @Param        id        	path      string  true  "Song ID"
// @Param        title     	form      string  true  "Updated song title"
// @Param        artist    	form      string  true  "Updated song artist"
// @Success      200            {object}  models.Song
// @Failure      400            "Bad Request"
// @Failure      404            "Song not found"
// @Failure      500            "Something went wrong"
// @Router       /collections/songs/{id} [put]

func UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	// Extract song ID from the URL path parameters
	songIDStr := chi.URLParam(r, "id")
	songID, err := uuid.FromString(songIDStr)
	if err != nil {
		logrus.Errorf("error parsing song ID: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the song with the given ID exists
	existingSong, err := collections.GetSongById(songID)
	if err != nil {
		logrus.Errorf("error fetching existing song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if existingSong == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Decode JSON request body into the 'song' struct
	var updatedSong models.Song
	err = json.NewDecoder(r.Body).Decode(&updatedSong)
	if err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request body"}`))
		return
	}

	// Validate the input data
	if updatedSong.Title == "" || updatedSong.Artist == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Title and artist are required fields"}`))
		return
	}

	// Call the repository function to update the song
	err = collections.UpdateSong(songID, updatedSong.Title, updatedSong.Artist)
	if err != nil {
		logrus.Errorf("error updating song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the updated song in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//responseBody, _ := json.Marshal(updatedSong)
	//_, _ = w.Write(responseBody)
	return
}
