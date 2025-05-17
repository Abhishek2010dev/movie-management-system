-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie_genre (
	movie_id INT  REFERENCES movie(id) ON DELETE CASCADE,
	genre_id INT REFERENCES genre(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movie_genre;
-- +goose StatementEnd
