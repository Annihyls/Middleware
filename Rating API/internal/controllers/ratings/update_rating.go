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
// @Summary      Update a rating.
// @Description  Update a rating.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id}  [put]
func UpdateRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("ratingId").(uuid.UUID)
	var rating models.Rating

    json.NewDecoder(r.Body).Decode(&rating)
	err := ratings.UpdateRating(ratingId, rating.Note,
	        rating.Description)

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
