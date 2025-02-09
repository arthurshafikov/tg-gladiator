-- +goose Up
-- +goose StatementBegin
CREATE TABLE hero_shop_items (
    hero_shop_id BIGINT NOT NULL,
    item_id BIGINT NOT NULL,
    price INT NOT NULL CHECK (price >= 0),

    FOREIGN KEY (hero_shop_id) REFERENCES hero_shops(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,

    PRIMARY KEY (hero_shop_id, item_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hero_shop_items;
-- +goose StatementEnd
