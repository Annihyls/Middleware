package collections

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	_ "middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
)

// DeleteSongHandler handles HTTP requests to delete an existing song from the collection.
// @Tags         collections
// @Summary      Delete an existing song from the collection.
// @Description  Delete an existing song from the collection.
// @Param        id        	path      string  true  "Song ID"
// @Success      204            "No Content"
// @Failure      400            "Bad Request"
// @Failure      404            "Song not found"
// @Failure      500            "Something went wrong"
// @Router       /collections/songs/{id} [delete]
func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
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

	// Call the repository function to delete the song
	err = collections.DeleteSong(songID)
	if err != nil {
		logrus.Errorf("error deleting song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusNoContent)
}
