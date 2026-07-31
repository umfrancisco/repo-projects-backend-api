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

type ProjectsResponse struct {
	Projects []*Project `json:"projects"`
}

// encore:api public method=GET path=/projects
func GetProjects(ctx context.Context) (*ProjectsResponse, error) {
	rows, err := DB.Query(ctx, `
		SELECT name, description, language, link
		FROM projects
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*Project

	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.Name, &p.Description, &p.Language, &p.Link); err != nil {
			return nil, err
		}
		projects = append(projects, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ProjectsResponse{
		Projects: projects,
	}, nil
}

// // encore:api public method=POST path=/project
// func CreateProject(ctx context.Context, p *Project) error {
// 	_, err := DB.Exec(ctx, `
// 		INSERT INTO projects (name, description, language, link)
// 		VALUES ($1, $2, $3, $4)
// 	`, p.Name, p.Description, p.Language, p.Link)

// 	return err
// }
