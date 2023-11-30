package ratings

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/ratings"
	"net/http"
)

// UpdateRating
// @Tags         collections
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Collection
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [get]
func CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
    json.NewDecoder(r.Body).Decode(&rating)

    //On créer un nouvel uuid car c'est plus simple que de devoir le renseigner !
    u, err := uuid.NewV4()
    if err != nil {
    	logrus.Errorf("failed to generate UUID: %s", err.Error())
    }

	err = ratings.CreateRating(u, rating.Note, rating.Description)

	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
