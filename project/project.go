package project

import "context"

// Project represents the response structure
type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Link        string `json:"link"`
}

// encore:api public method=GET path=/project
func GetProject(ctx context.Context) (*Project, error) {
	return &Project{
		Name:        "project-name",
		Description: "short-desc",
		Language:    "java",
		Link:        "link-to-github-repo",
	}, nil
}