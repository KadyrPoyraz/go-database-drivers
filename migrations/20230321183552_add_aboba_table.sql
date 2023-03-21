-- +goose Up
CREATE TABLE aboba (
    id SERIAL NOT NULL,
    name VARCHAR(255),
    weight INTEGER,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE aboba;
