CREATE TABLE users
(
    id   SERIAL       NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE ,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE chat
(
    id SERIAL NOT NULL PRIMARY KEY ,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE message
(
    id SERIAL NOT NULL PRIMARY KEY ,
    user_id INT REFERENCES users(id) NOT NULL ,
    chat_id INT REFERENCES chat(id) NOT NULL ,
    date TIMESTAMP NOT NULL ,
    text TEXT NOT NULL
);

CREATE VIEW message_data AS
    SELECT u.name AS name, m.date AS date,
           m.text AS text, m.chat_id AS chat_id,
           m.id AS id
    FROM message m
    JOIN users u on m.user_id = u.id;
