-- +goose Up
CREATE TABLE test_table (
    id INT NOT NULL,
    name VARCHAR(255),
    weight INTEGER,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE test_table;
