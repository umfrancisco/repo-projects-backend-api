CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    language TEXT NOT NULL,
    link TEXT NOT NULL
);