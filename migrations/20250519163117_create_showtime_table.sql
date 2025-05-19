-- +goose Up
-- +goose StatementBegin
CREATE TABLE showtime (
	id SERIAL PRIMARY KEY,
	movie_id INT NOT NULL REFERENCES movie(id) ON DELETE CASCADE,
	start_time TIMESTAMP NOT NULL,
	end_time TIMESTAMP NOT NULL,
	available_seats INT NOT NULL,
	price NUMERIC(10, 2) NOT NULL, 
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE showtime;
-- +goose StatementEnd
