package users

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/users"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateUser
// @Tags         users
// @Summary      Get a user.
// @Description  Get a user.
// @Param        id           	path      string  true  "User UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [get]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	err := users.UpdateUser(userId, user.Note, user.Description)

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
