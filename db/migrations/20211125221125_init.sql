-- migrate:up

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;



-------- User 
CREATE TABLE IF NOT EXISTS users (
  -- default
  user_id       BIGSERIAL         PRIMARY KEY,
  created_at    TIMESTAMP         NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  -- main 
  username      VARCHAR(80)       NOT NULL UNIQUE,
  email         VARCHAR(80)       NOT NULL UNIQUE,
  password      VARCHAR(700)      NOT NULL
);
-- when the row is updated, update the "updated_at" timestamp
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();



-------- Post
CREATE TABLE IF NOT EXISTS posts (
  -- default
  post_id       BIGSERIAL       PRIMARY KEY,
  created_at    TIMESTAMP       NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
  -- main 
  post_title    VARCHAR(80)     NOT NULL,
  post_body     VARCHAR(500)    NOT NULL,
  -- relationships
  user_id       int             NOT NULL REFERENCES users(user_id) ON DELETE CASCADE
);
-- when the row is updated, update the "updated_at" timestamp
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON posts
FOR EACH ROW 
EXECUTE PROCEDURE trigger_set_timestamp();



-- create fake users
INSERT INTO users ( email, username, password ) VALUES 
( 'user1@email.com', 'random-user-1', 'not-hashed-password' ),
( 'user2@email.com', 'random-user-2', 'not-hashed-password' ),
( 'user3@email.com', 'random-user-3', 'not-hashed-password' ),
( 'user4@email.com', 'random-user-4', 'not-hashed-password' ),
( 'user5@email.com', 'random-user-5', 'not-hashed-password' ),
( 'user6@email.com', 'random-user-6', 'not-hashed-password' ),
( 'user7@email.com', 'random-user-7', 'not-hashed-password' ),
( 'user8@email.com', 'random-user-8', 'not-hashed-password' ),
( 'user9@email.com', 'random-user-9', 'not-hashed-password' ),
( 'user10@email.com','random-user-10','not-hashed-password' );



-- create fake posts
INSERT INTO posts ( post_title, post_body, user_id  ) VALUES 
( 'this is the title', 'this is the body of my post', 2 ),
( 'this is the title', 'this is the body of my post', 2 ),
( 'this is the title', 'this is the body of my post', 2 ),
( 'this is the title', 'this is the body of my post', 1 ),
( 'this is the title', 'this is the body of my post', 1 ),
( 'this is the title', 'this is the body of my post', 1 ),
( 'this is the title', 'this is the body of my post', 3 ),
( 'this is the title', 'this is the body of my post', 3 );



-- migrate:down

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;

