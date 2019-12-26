CREATE TABLE users
(
    id SERIAL,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
