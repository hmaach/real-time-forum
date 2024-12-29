package models

import (
	"fmt"
	"net/http"
	"time"
)

func StoreSession(user_id int, session_id string, expires_at time.Time) error {
	query := `INSERT OR REPLACE INTO sessions (user_id,session_id,expires_at) VALUES (?,?,?)`

	_, err := DB.Exec(query, user_id, session_id, expires_at)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func ValidSession(r *http.Request) (int, string, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		return -1, "", false
	}
	var expiration time.Time
	var user_id int
	var username string
	query := `
		SELECT 
			s.user_id,
			s.expires_at, 
			u.username 
		FROM sessions s 
		INNER JOIN users u ON s.user_id = u.id 
		WHERE session_id = ?
	`
	err = DB.QueryRow(query, cookie.Value).Scan(&user_id, &expiration, &username)
	if err != nil || expiration.Before(time.Now()) {
		return -1, "", false
	}
	return user_id, username, true
}

func DeleteUserSession(userID int) error {
	_, err := DB.Exec(`DELETE FROM sessions WHERE user_id = ?;`, userID)
	return err
}
