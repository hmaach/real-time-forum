package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"

	"forum/server/models"
	"forum/server/utils"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetLoginPage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}

	var valid bool
	if _, _, valid = models.ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err := utils.RenderTemplate(db, w, r, "login", http.StatusOK, nil, false, "")
	if err != nil {
		log.Println(err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, false, "")
		return
	}
}

func Signin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}

	var valid bool
	if _, _, valid = models.ValidSession(r, db); valid {
		w.WriteHeader(302)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(strings.TrimSpace(username)) < 4 || len(strings.TrimSpace(password)) < 6 {
		w.WriteHeader(400)
		return
	}

	// get user information from database
	user_id, hashedPassword, err := models.GetUserInfo(db, username)
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
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		log.Println("Failed to create session")
		return
	}
	sessionID := sessionId.String()

	err = models.StoreSession(db, user_id, sessionID, time.Now().Add(10*time.Hour))
	if err != nil {
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		log.Println("Failed to create session")
		return
	}

	// Set session ID as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userID, username, valid := models.ValidSession(r, db)

	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, valid, username)
		return
	}

	if valid {
		// Use the new model function
		err := models.DeleteUserSession(db, userID)
		if err != nil {
			utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
