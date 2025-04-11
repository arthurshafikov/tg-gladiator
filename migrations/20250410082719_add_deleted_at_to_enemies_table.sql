-- +goose Up
-- +goose StatementBegin
ALTER TABLE enemies
ADD COLUMN deleted_at TIMESTAMP NULL;

UPDATE enemies
SET deleted_at = NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE enemies
DROP COLUMN deleted_at;
-- +goose StatementEnd
