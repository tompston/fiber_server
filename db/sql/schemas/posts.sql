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
