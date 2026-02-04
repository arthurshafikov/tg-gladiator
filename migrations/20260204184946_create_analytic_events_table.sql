
-- +goose Up
-- +goose StatementBegin
CREATE TABLE analytic_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chat_id BIGINT NULL,
    type VARCHAR(64) NOT NULL,
    payload JSONB NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_analytic_events_chat FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE SET NULL
);
CREATE INDEX idx_analytic_events_timestamp ON analytic_events (timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_analytic_events_timestamp;
DROP TABLE IF EXISTS analytic_events;
-- +goose StatementEnd
