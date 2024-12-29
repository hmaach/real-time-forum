package models

import (
	"fmt"
)

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	UserName  string
	Content   string
	Likes     int
	Dislikes  int
	CreatedAt string
}

func FetchCommentsByPostID(postID int) ([]Comment, error) {
	var comments []Comment
	query := `
	SELECT
		c.id,
		c.user_id,
		u.username,
		c.content,
		strftime('%m/%d/%Y %I:%M %p', c.created_at) AS formatted_created_at,
		(
			SELECT
				COUNT(*)
			FROM
				comment_reactions AS cr
			WHERE
				cr.comment_id = c.id
				AND cr.reaction = 'like'
		) AS likes_count,
		(
			SELECT
				COUNT(*)
			FROM
				comment_reactions AS cr
			WHERE
				cr.comment_id = c.id
				AND cr.reaction = 'dislike'
		) AS dislikes_count
	FROM
		comments c
	INNER JOIN users u 
	ON c.user_id = u.id
	WHERE
		c.post_id = ?
	ORDER BY
		c.created_at DESC
	`

	rows, err := DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.UserName,
			&comment.Content,
			&comment.CreatedAt,
			&comment.Likes,
			&comment.Dislikes,
		)
		if err != nil {
			return nil, err
		}

		// Assign the post ID and format the created_at field
		comment.PostID = postID
		// comment.CreatedAt = utils.FormatTime(comment.CreatedAt)

		// Append the comment to the slice
		comments = append(comments, comment)
	}

	return comments, nil
}

func StoreComment(user_id, post_id int, content string) (int64, error) {
	query := `INSERT INTO comments (user_id,post_id,content) VALUES (?,?,?)`

	result, err := DB.Exec(query, user_id, post_id, content)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	commentID, _ := result.LastInsertId()

	return commentID, nil
}

func StoreCommentReaction(user_id, comment_id int, reaction string) (int64, error) {
	query := `INSERT INTO comment_reactions (user_id,comment_id,reaction) VALUES (?,?,?)`
	result, err := DB.Exec(query, user_id, comment_id, reaction)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("error inserting reaction data -> ")
	}
	creactionID, _ := result.LastInsertId()

	return creactionID, nil
}

// Count comments by post ID
func CountCommentsByPostID(postID int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM comments WHERE post_id = ?"
	err := DB.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error counting comments: %v", err)
	}
	return count, nil
}

// Fetch the creation time of a comment by its ID
func FetchCommentTimeByID(commentID int64) (string, error) {
	var commentTime string
	query := "SELECT strftime('%m/%d/%Y %I:%M %p', created_at) AS formatted_created_at FROM comments WHERE id = ?"
	err := DB.QueryRow(query, commentID).Scan(&commentTime)
	if err != nil {
		return "", fmt.Errorf("error fetching comment time: %v", err)
	}
	return commentTime, nil
}

func ReactToComment(user_id, comment_id int, userReaction string) (int, int, error) {
	var likeCount, dislikeCount int
	var dbreaction string
	var err error

	DB.QueryRow("SELECT reaction FROM comment_reactions WHERE user_id=? AND comment_id=?", user_id, comment_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = StoreCommentReaction(user_id, comment_id, userReaction)
	} else {
		if userReaction == dbreaction {
			query := "DELETE FROM comment_reactions WHERE user_id = ? AND comment_id = ?"
			_, err = DB.Exec(query, user_id, comment_id)

		} else {
			query := "UPDATE comment_reactions SET reaction = ? WHERE user_id = ? AND comment_id = ?"
			_, err = DB.Exec(query, userReaction, user_id, comment_id)
		}
	}
	if err != nil {
		return 0, 0, err
	}

	// Fetch the new count of reactions for this post
	err = DB.QueryRow("SELECT COUNT(*) FROM comment_reactions WHERE comment_id=? AND reaction=?", comment_id, "like").Scan(&likeCount)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching likes count: %v", err)
	}
	err = DB.QueryRow("SELECT COUNT(*) FROM comment_reactions WHERE comment_id=? AND reaction=?", comment_id, "dislike").Scan(&dislikeCount)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching likes count: %v", err)
	}
	return likeCount, dislikeCount, nil
}
