package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/ratings"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/ratings", func(r chi.Router) {
		r.Get("/", ratings.GetRatings)
		r.Post("/", ratings.CreateRating)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(ratings.Ctx)
			r.Put("/", ratings.UpdateRating)
			r.Delete("/", ratings.DeleteRating)
			r.Get("/", ratings.GetRating)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8079")
	logrus.Fatalln(http.ListenAndServe(":8079", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS ratings (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			note INTEGER NOT NULL,
			description VARCHAR(255),
			id_user VARCHAR(255) NOT NULL,
			id_song VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
