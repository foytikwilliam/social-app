DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS ratings;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS threads;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS gyms;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  name       VARCHAR(255),
  email      VARCHAR(255) NOT NULL UNIQUE,
  password   VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE gyms (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  name       VARCHAR(255),
  address    VARCHAR(255),
  city       VARCHAR(255),
  state      VARCHAR(255),
  zipcode    VARCHAR(20),
  county     VARCHAR(255),
  phone      VARCHAR(20),
  email      VARCHAR(255),
  website    VARCHAR(255),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE sessions (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  email      VARCHAR(255),
  user_id    INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE threads (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  topic      TEXT,
  user_id    INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE posts (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  body       TEXT,
  user_id    INTEGER REFERENCES users(id),
  thread_id  INTEGER REFERENCES threads(id),
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE reviews (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  body       TEXT,
  user_id    INTEGER REFERENCES users(id),
  gym_id     INTEGER REFERENCES gyms(id),
  rating     INTEGER,
  date       TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE ratings (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  user_id    INTEGER REFERENCES users(id),
  gym_id     INTEGER REFERENCES gyms(id),
  rating     INTEGER,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE likes (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  user_id    INTEGER REFERENCES users(id),
  gym_id     INTEGER REFERENCES gyms(id),
  post_id    INTEGER REFERENCES posts(id),
  created_at TIMESTAMP NOT NULL
);
