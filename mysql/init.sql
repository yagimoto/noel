USE noel-db;

CREATE TABLE users (
    user_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE token (
    token_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    token VARCHAR(255) NOT NULL,
    user_id INT NOT NULL, 
    -- VARCHARとなっていた
    expiration_time TIMESTAMP NOT NULL
);

CREATE TABLE scores (
    score_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    score INT NOT NULL,
    user_id INT NOT NULL
);

CREATE TABLE cities (
    city_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    city_name VARCHAR(255) NOT NULL,
    score int NOT NULL
);

-- 以下デバッグ用

INSERT into scores (score_id, score, user_id) VALUES(1, 234, 1);

INSERT into scores (score_id, score, user_id) VALUES(2, 8, 2);

INSERT into scores (score_id, score, user_id) VALUES(3, 520, 3);

INSERT into cities (city_id, city_name, score) VALUES(73, "TOKYO", 611);

INSERT into cities (city_id, city_name, score) VALUES(2, "OSAKA", 411);

INSERT into cities (city_id, city_name, score) VALUES(5, "OKINAWA", 511);