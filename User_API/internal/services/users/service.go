package users

import (
	"database/sql"
	"errors"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/users"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}
func UpdateUser(id uuid.UUID, prenom string, nom string) error {
	err := repository.UpdateUser(id, prenom, nom)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating user : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return err
}

func DeleteUser(id uuid.UUID) error {
	err := repository.DeleteUser(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating user : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return err
}

func CreateUser(id uuid.UUID, prenom string, nom string) error {
	err := repository.CreateUser(id, prenom, nom)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating user : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return err
}
