package project

import (
	"context"
)

// Project represents the response structure
type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Link        string `json:"link"`
}

// encore:api public method=GET path=/project
func GetProject(ctx context.Context) (*Project, error) {
	row := DB.QueryRow(ctx, `
		SELECT name, description, language, link
		FROM projects
	`)

	var p Project
	err := row.Scan(&p.Name, &p.Description, &p.Language, &p.Link)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// encore:api public method=POST path=/project
func CreateProject(ctx context.Context, p *Project) error {
	_, err := DB.Exec(ctx, `
		INSERT INTO projects (name, description, language, link)
		VALUES ($1, $2, $3, $4)
	`, p.Name, p.Description, p.Language, p.Link)

	return err
}
