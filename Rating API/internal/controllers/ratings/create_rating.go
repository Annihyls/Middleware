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
// @Tags         ratings
// @Summary      Create a rating.
// @Description  Create a rating. UUID is automatically generated
// @Success      201            {object}  models.Rating
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /ratings       [post]
func CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
    json.NewDecoder(r.Body).Decode(&rating)

    //On créer un nouvel uuid car c'est plus simple que de devoir le renseigner !
    u, err := uuid.NewV4()
    rating.Id = &u
    if err != nil {
    	logrus.Errorf("failed to generate UUID: %s", err.Error())
    }

	err = ratings.CreateRating(*rating.Id, rating.Note, rating.Description)

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

    w.Header().Set("Location", "/ratings/" + rating.Id.String())
	w.WriteHeader(http.StatusCreated)
    body, _ := json.Marshal(rating)
    _, _ = w.Write(body)
	return
}