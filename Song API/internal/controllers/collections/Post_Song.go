package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostSongHandler handles HTTP requests to add a new song to the collection.
// @Tags         collections
// @Summary      Add a new song to the collection.
// @Description  Add a new song to the collection.
// @Param        title        	form      string  true  "Song title"
// @Param        artist       	form      string  true  "Song artist"
// @Success      201            {object}  models.Song
// @Failure      400            "Bad Request"
// @Failure      500            "Something went wrong"
// @Router       /collections/songs [post]

func PostSongHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.FormValue("title")
	//artist := r.FormValue("artist")
	var song models.Song
	json.NewDecoder(r.Body).Decode(&song)
	// Validate the input data

	newSong, err := collections.PostSong(song.Title, song.Artist)
	if err != nil {
		logrus.Errorf("error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", "/collections/songs/"+newSong.ID.String())
	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(newSong)
	_, _ = w.Write(body)
	return
}
