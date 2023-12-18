package users

import (
	"encoding/json"
	"git2/Middleware/User_API/internal/models"
	"git2/Middleware/User_API/internal/repositories/users"
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
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)

	err := users.DeleteUser(userId)

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
