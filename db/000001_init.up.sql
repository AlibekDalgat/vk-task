CREATE TABLE users
(
    login VARCHAR(255) PRIMARY KEY,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE advertisements
(
    id SERIAL PRIMARY KEY ,
    title VARCHAR(255) NOT NULL,
    text VARCHAR(1000),
    image VARCHAR(255),
    price REAL,
    owner VARCHAR(255) REFERENCES users(login) ON DELETE CASCADE NOT NULL
);