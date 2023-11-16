--+migrate Up 
CREATE TABLE IF NOT EXISTS "user" (
  id      SERIAL PRIMARY KEY,
  name    VARCHAR(100),
  email   VARCHAR(50) UNIQUE,
  api_key TEXT UNIQUE,
  permission_level INT DEFAULT 1,
);

--+migrate Down 
DROP TABLE "user";
