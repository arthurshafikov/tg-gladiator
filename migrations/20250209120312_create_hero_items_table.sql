-- +goose Up
-- +goose StatementBegin
CREATE TABLE hero_items (
    hero_id BIGINT NOT NULL,
    item_id BIGINT NOT NULL,
    is_equipped BOOLEAN NOT NULL DEFAULT(false),
    quantity INT NOT NULL CHECK (quantity > 0),

    FOREIGN KEY (hero_id) REFERENCES heroes(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,

    PRIMARY KEY (hero_id, item_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hero_items;
-- +goose StatementEnd
