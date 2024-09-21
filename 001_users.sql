-- +goose up
CREATE TABLE users (
    id int NOT NULL PRIMARY KEY,
    username text,
    name text,
    surname text
);

-- +goose statementbegin
INSERT INTO "users" ("id", "username") VALUES (1, 'gallant_almeida7');
INSERT INTO "users" ("id", "username") VALUES (2, 'brave_spence8');
.
.
INSERT INTO "users" ("id", "username") VALUES (99999, 'jovial_chaum1');
INSERT INTO "users" ("id", "username") VALUES (100000, 'goofy_ptolemy0');
-- +goose statementend

-- +goose down
DROP TABLE users;