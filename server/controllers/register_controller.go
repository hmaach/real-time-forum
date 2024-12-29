package controllers

import (
	"net/http"

	"forum/server/models"
	"forum/server/validators"
)

func Register(w http.ResponseWriter, r *http.Request) {
	statusCode, _, email, username, password := validators.RegisterRequest(r)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return
	}

	var valid bool
	if _, _, valid = models.ValidSession(r); valid {
		w.WriteHeader(302)
		return
	}

	_, err := models.StoreUser(email, username, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" || err.Error() == "UNIQUE constraint failed: users.email" {
			w.WriteHeader(304)
			return
		}

		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
