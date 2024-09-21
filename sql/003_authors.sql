-- +goose up
CREATE TABLE authors
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    bio  TEXT NOT NULL
);


-- +goose statementbegin
INSERT INTO "authors" ("name","bio") VALUES ( 'edgar_alan_poe', 'great writer');
INSERT INTO "authors" ("name","bio") VALUES ( 'sidney_jackson', 'great writer');
-- +goose statementend

-- +goose down
DROP TABLE authors;