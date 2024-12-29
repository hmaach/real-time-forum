package models

import (
	"fmt"
	"strings"
)

type Category struct {
	ID         int
	Label      string
	PostsCount int
}

func FetchCategories() ([]Category, error) {
	var categories []Category
	query := `
		SELECT
			c.id,
			c.label,
			(
				SELECT
					COUNT(id)
				FROM
					post_category pc
				WHERE
					pc.category_id = c.id
			) as posts_count
		FROM categories c
		ORDER BY posts_count DESC;
	`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		rows.Scan(&category.ID, &category.Label, &category.PostsCount)
		categories = append(categories, category)
	}
	return categories, nil
}

func CheckCategories(ids []int) error {
	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1]

	query := fmt.Sprintf(`
        SELECT id
        FROM categories
        WHERE id IN (%s);
    `, placeholders)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return err
		}
		count++
	}
	if count != len(ids) {
		return fmt.Errorf("categories does not exists in db")
	}

	return nil
}
