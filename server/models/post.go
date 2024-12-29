package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Post struct {
	ID            int
	UserID        int
	UserName      string
	Title         string
	Content       string
	CreatedAt     string
	Likes         int
	Dislikes      int
	Comments      int
	CategoriesStr string
	Categories    []string
}

type PostDetail struct {
	Post     Post
	Comments []Comment
}

func FetchPosts(page int) ([]Post, error) {
	var posts []Post
	// Query to fetch posts
	query := `SELECT
		p.id,
		p.user_id,
		u.username,
		p.title,
		p.content,
		strftime('%m/%d/%Y %I:%M %p', p.created_at) AS formatted_created_at,
		(
			SELECT
				COUNT(*)
			FROM
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'like'
		) AS likes_count,
		(
			SELECT
				COUNT(*)
			FROM
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'dislike'
		) AS dislikes_count,
		(
			SELECT
				COUNT(*)
			FROM
				comments c
			WHERE
				c.post_id = p.id
		) AS comments_count,
		(
			SELECT
				GROUP_CONCAT(c.label)
			FROM
				categories c
			INNER JOIN post_category pc ON c.id = pc.category_id
			WHERE
				pc.post_id = p.id
		) AS categories
	FROM
		posts p
		INNER JOIN users u ON p.user_id = u.id
	ORDER BY
		p.created_at DESC
	LIMIT 10 OFFSET ? ;
	`
	rows, err := DB.Query(query, page)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var post Post
		// Scan the data into the Post struct
		err := rows.Scan(&post.ID,
			&post.UserID,
			&post.UserName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&post.CategoriesStr)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		// it came from the  database as "technology,sports...", so we need to split it
		post.Categories = strings.Split(post.CategoriesStr, ",")

		// Append the Post struct to the posts slice
		posts = append(posts, post)
	}
	return posts, nil
}

func FetchPost(postID int) (PostDetail, error) {
	var post Post
	post.ID = postID

	// Query to fetch the post
	query := `SELECT
		p.user_id,
		u.username,
		p.title,
		p.content,
		strftime('%m/%d/%Y %I:%M %p', p.created_at) AS formatted_created_at,
		(
			SELECT COUNT(*)
			FROM post_reactions AS pr
			WHERE pr.post_id = p.id
			AND pr.reaction = 'like'
		) AS likes_count,
		(
			SELECT COUNT(*)
			FROM post_reactions AS pr
			WHERE pr.post_id = p.id
			AND pr.reaction = 'dislike'
		) AS dislikes_count,
		(
			SELECT COUNT(*)
			FROM comments c
			WHERE c.post_id = p.id
		) AS comments_count,
		(
			SELECT GROUP_CONCAT(c.label)
			FROM categories c
			INNER JOIN post_category pc ON c.id = pc.category_id
			WHERE pc.post_id = p.id
		) AS categories
	FROM
		posts p
		INNER JOIN users u ON p.user_id = u.id
	WHERE p.id = ?`

	// Use QueryRow for a single result
	row := DB.QueryRow(query, postID)

	// Scan the data into the Post struct
	err := row.Scan(
		&post.UserID,
		&post.UserName,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.Likes,
		&post.Dislikes,
		&post.Comments,
		&post.CategoriesStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return PostDetail{}, err
		}
		log.Println("Error scanning row:", err)
		return PostDetail{}, err
	}

	// Process categories
	post.Categories = strings.Split(post.CategoriesStr, ",")

	// Format the created_at field
	// post.CreatedAt = post.CreatedAt.Format("01/02/2006 03:04 PM")
	comments, err := FetchCommentsByPostID(postID)
	if err != nil {
		log.Println("Error fetching comments from the database:", err)
	}

	return PostDetail{
		Post:     post,
		Comments: comments,
	}, nil
}

func FetchPostsByCategory(categoryID, page int) ([]Post, error) {
	var posts []Post
	query := `
		SELECT
			p.id,
			p.user_id,
			u.username,
			p.title,
			p.content,
			strftime('%m/%d/%Y %I:%M %p', p.created_at) AS formatted_created_at,
			(
				SELECT
					COUNT(*)
				FROM
					post_reactions AS pr
				WHERE
					pr.post_id = p.id
					AND pr.reaction = 'like'
			) AS likes_count,
			(
				SELECT
					COUNT(*)
				FROM
					post_reactions AS pr
				WHERE
					pr.post_id = p.id
					AND pr.reaction = 'dislike'
			) AS dislikes_count,
			(
				SELECT
					COUNT(*)
				FROM
					comments c
				WHERE
					c.post_id = p.id
			) AS comments_count,
			(
				SELECT
					GROUP_CONCAT(c.label)
				FROM
					categories c
				INNER JOIN post_category pc ON c.id = pc.category_id
				WHERE
					pc.post_id = p.id
			) AS categories
		FROM
			posts p
			INNER JOIN users u ON p.user_id = u.id
			INNER JOIN post_category pc ON p.id = pc.post_id
		WHERE pc.category_id = ?
		ORDER BY
			p.created_at
		LIMIT 10 OFFSET ? ;
	`
	rows, err := DB.Query(query, categoryID, page)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID,
			&post.UserID,
			&post.UserName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&post.CategoriesStr)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		// it came from the  database as "technology,sports...", so we need to split it
		post.Categories = strings.Split(post.CategoriesStr, ",")

		// post.CreatedAt = utils.FormatTime(post.CreatedAt)

		posts = append(posts, post)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return posts, nil
}

func FetchCreatedPostsByUser(user_id int) ([]Post, error) {
	var posts []Post
	currentPage := 0
	// Query to fetch posts
	query := `SELECT
		p.id,
		p.user_id,
		u.username,
		p.title,
		p.content,
		strftime('%m/%d/%Y %I:%M %p', p.created_at) AS formatted_created_at,
		(
			SELECT
				COUNT(*)
			FROM
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'like'
		) AS likes_count,
		(
			SELECT
				COUNT(*)
			FROM
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'dislike'
		) AS dislikes_count,
		(
			SELECT
				COUNT(*)
			FROM
				comments c
			WHERE
				c.post_id = p.id
		) AS comments_count,
		(
			SELECT
				GROUP_CONCAT(c.label)
			FROM
				categories c
			INNER JOIN post_category pc ON c.id = pc.category_id
			WHERE
				pc.post_id = p.id
		) AS categories
	FROM
		posts p
		INNER JOIN users u ON p.user_id = u.id
	WHERE p.user_id = ?
	ORDER BY
		p.created_at DESC
	LIMIT 10 OFFSET ? ;
	`
	rows, err := DB.Query(query, user_id, currentPage)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var post Post
		// Scan the data into the Post struct
		err := rows.Scan(&post.ID,
			&post.UserID,
			&post.UserName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&post.CategoriesStr)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		// it came from the  database as "technology,sports...", so we need to split it
		post.Categories = strings.Split(post.CategoriesStr, ",")

		// Format the created_at field to a more readable format
		// post.CreatedAt = utils.FormatTime(post.CreatedAt)

		// Append the Post struct to the posts slice
		posts = append(posts, post)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return posts, nil
}

func FetchLikedPostsByUser(user_id int) ([]Post, error) {
	var posts []Post
	currentPage := 0
	// Query to fetch posts
	query := `SELECT
		p.id,
		p.user_id,
		u.username,
		p.title,
		p.content,
		strftime('%m/%d/%Y %I:%M %p', p.created_at) AS formatted_created_at,
		(
			SELECT
				COUNT(*)
			FROM
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'like'
		) AS likes_count,
		(
			SELECT
				COUNT(*)
			FROM
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'dislike'
		) AS dislikes_count,
		(
			SELECT
				COUNT(*)
			FROM
				comments c
			WHERE
				c.post_id = p.id
		) AS comments_count,
		(
			SELECT
				GROUP_CONCAT(c.label)
			FROM
				categories c
			INNER JOIN post_category pc ON c.id = pc.category_id
			WHERE
				pc.post_id = p.id
		) AS categories
	FROM
		posts p
		INNER JOIN users u ON p.user_id = u.id
		INNER JOIN post_reactions pr ON p.id = pr.post_id
	WHERE pr.user_id = ? AND pr.reaction = 'like' 
	ORDER BY
		p.created_at DESC
	LIMIT 10 OFFSET ? ;
	`
	rows, err := DB.Query(query, user_id, currentPage)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var post Post
		// Scan the data into the Post struct
		err := rows.Scan(&post.ID,
			&post.UserID,
			&post.UserName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&post.CategoriesStr)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		// it came from the  database as "technology,sports...", so we need to split it
		post.Categories = strings.Split(post.CategoriesStr, ",")

		// Format the created_at field to a more readable format
		// post.CreatedAt = utils.FormatTime(post.CreatedAt)

		// Append the Post struct to the posts slice
		posts = append(posts, post)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return posts, nil
}

func StorePost(user_id int, title, content string) (int64, error) {
	query := `INSERT INTO posts (user_id,title,content) VALUES (?,?,?)`

	result, err := DB.Exec(query, user_id, title, content)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	postID, _ := result.LastInsertId()

	return postID, nil
}

func StorePostCategory(post_id int64, category_id int) (int64, error) {
	query := `INSERT INTO post_category (post_id, category_id) VALUES (?,?)`

	result, err := DB.Exec(query, post_id, category_id)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	postcatID, _ := result.LastInsertId()

	return postcatID, nil
}

func StorePostReaction(user_id, post_id int, reaction string) (int64, error) {
	query := `INSERT INTO post_reactions (user_id,post_id,reaction) VALUES (?,?,?)`
	result, err := DB.Exec(query, user_id, post_id, reaction)
	if err != nil {
		return 0, fmt.Errorf("error inserting reaction data -> ")
	}
	preactionID, _ := result.LastInsertId()

	return preactionID, nil
}

func ReactToPost(user_id, post_id int, userReaction string) (int, int, error) {
	var likeCount, dislikeCount int
	var dbreaction string
	var err error
	DB.QueryRow("SELECT reaction FROM post_reactions WHERE user_id=? AND post_id=?", user_id, post_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = StorePostReaction(user_id, post_id, userReaction)
	} else {
		if userReaction == dbreaction {
			query := "DELETE FROM post_reactions WHERE user_id = ? AND post_id = ?"
			_, err = DB.Exec(query, user_id, post_id)
		} else {
			query := "UPDATE post_reactions SET reaction = ? WHERE user_id = ? AND post_id = ?"
			_, err = DB.Exec(query, userReaction, user_id, post_id)
		}
	}

	if err != nil {
		return 0, 0, err
	}

	// Fetch the new count of reactions for this post
	DB.QueryRow("SELECT COUNT(*) FROM post_reactions WHERE post_id=? AND reaction=?", post_id, "like").Scan(&likeCount)
	DB.QueryRow("SELECT COUNT(*) FROM post_reactions WHERE post_id=? AND reaction=?", post_id, "dislike").Scan(&dislikeCount)

	return likeCount, dislikeCount, nil
}
