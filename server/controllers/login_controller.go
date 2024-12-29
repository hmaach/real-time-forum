package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"forum/server/models"
	"forum/server/validators"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	statusCode, _, username, password := validators.LoginRequest(r)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return
	}
	if _, _, valid := models.ValidSession(r); valid {
		w.WriteHeader(302)
		return
	}

	// get user information from database
	user_id, hashedPassword, err := models.GetUserInfo(username)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		w.WriteHeader(401)
		return
	}

	sessionId, err := uuid.NewV7()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to create session")
		return
	}

	err = models.StoreSession(user_id, sessionId.String(), time.Now().Add(10*time.Hour))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to create session")
		return
	}

	// Set session ID as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionId.String(),
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userID, _, valid := models.ValidSession(r)

	if valid {
		err := models.DeleteUserSession(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error while logging out!")
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Expires:  time.Now(),
			HttpOnly: true,
			Path:     "/",
		})
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
}
