-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie (
    id SERIAL PRIMARY KEY,
    title VARCHAR(254) NOT NULL,
    description TEXT NOT NULL,
    release_date DATE NOT NULL,
    duration_minutes INT NOT NULL,
    director VARCHAR(100) NOT NULL,
    poster_path VARCHAR(254) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movie;
-- +goose StatementEnd
