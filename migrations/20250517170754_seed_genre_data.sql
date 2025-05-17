-- +goose Up
-- +goose StatementBegin
INSERT INTO genre (name, description) VALUES
('Action', 'Movies with intense physical activity and stunts'),
('Comedy', 'Movies designed to make the audience laugh'),
('Drama', 'Movies with emotionally intense storytelling'),
('Fantasy', 'Movies featuring magical or supernatural elements'),
('Horror', 'Movies designed to frighten and create suspense'),
('Romance', 'Movies focused on love stories'),
('Sci-Fi', 'Science fiction films with futuristic settings or technology'),
('Thriller', 'Movies full of tension, excitement, and mystery'),
('Documentary', 'Non-fiction films based on real events or people'),
('Animation', 'Movies created using animation techniques');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM genre;
-- +goose StatementEnd
