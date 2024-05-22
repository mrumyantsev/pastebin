-- +goose Up
CREATE TABLE users
(
	id            SERIAL       NOT NULL UNIQUE,
	username      VARCHAR(64)  NOT NULL UNIQUE,
	password_hash VARCHAR(256) NOT NULL,
	first_name    VARCHAR(256) NOT NULL,
	last_name     VARCHAR(256) NOT NULL,
	email         VARCHAR(256) NOT NULL UNIQUE,
	created_at    TIMESTAMPTZ  NOT NULL,
		CONSTRAINT pk_users PRIMARY KEY (id)
);

INSERT INTO users (username, password_hash, first_name, last_name, email, created_at)
VALUES
('asd1', '1234', 'fn', 'ln', 'e1@mail.com', '2024-05-19T12:25:11'),
('asd2', '1234', 'fn', 'ln', 'e2@mail.com', '2024-05-19T12:25:11'),
('asd3', '1234', 'fn', 'ln', 'e3@mail.com', '2024-05-19T12:25:11'),
('asd4', '1234', 'fn', 'ln', 'e4@mail.com', '2024-05-19T12:25:11'),
('asd5', '1234', 'fn', 'ln', 'e5@mail.com', '2024-05-19T12:25:11'),
('asd6', '1234', 'fn', 'ln', 'e6@mail.com', '2024-05-19T12:25:11'),
('asd7', '1234', 'fn', 'ln', 'e7@mail.com', '2024-05-19T12:25:11'),
('asd8', '1234', 'fn', 'ln', 'e8@mail.com', '2024-05-19T12:25:11'),
('asd9', '1234', 'fn', 'ln', 'e9@mail.com', '2024-05-19T12:25:11'),
('asd10', '1234', 'fn', 'ln', 'e10@mail.com', '2024-05-19T12:25:11'),
('asd11', '1234', 'fn', 'ln', 'e11@mail.com', '2024-05-19T12:25:11'),
('asd12', '1234', 'fn', 'ln', 'e12@mail.com', '2024-05-19T12:25:11'),
('asd13', '1234', 'fn', 'ln', 'e13@mail.com', '2024-05-19T12:25:11'),
('asd14', '1234', 'fn', 'ln', 'e14@mail.com', '2024-05-19T12:25:11'),
('asd15', '1234', 'fn', 'ln', 'e15@mail.com', '2024-05-19T12:25:11'),
('asd16', '1234', 'fn', 'ln', 'e16@mail.com', '2024-05-19T12:25:11'),
('asd17', '1234', 'fn', 'ln', 'e17@mail.com', '2024-05-19T12:25:11'),
('asd18', '1234', 'fn', 'ln', 'e18@mail.com', '2024-05-19T12:25:11'),
('asd19', '1234', 'fn', 'ln', 'e19@mail.com', '2024-05-19T12:25:11'),
('asd20', '1234', 'fn', 'ln', 'e20@mail.com', '2024-05-19T12:25:11'),
('asd21', '1234', 'fn', 'ln', 'e21@mail.com', '2024-05-19T12:25:11');

-- +goose Down
DROP TABLE users;
