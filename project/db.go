package project

import "encore.dev/storage/sqldb"

// Create a PostgreSQL database
var DB = sqldb.NewDatabase("project_db", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
