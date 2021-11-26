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