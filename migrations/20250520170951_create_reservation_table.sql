-- +goose Up
-- +goose StatementBegin
CREATE TABLE reservation (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    showtime_id INT NOT NULL REFERENCES showtime(id) ON DELETE CASCADE,
    reservation_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reservation;
-- +goose StatementEnd
