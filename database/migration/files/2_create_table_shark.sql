-- +migrate Up 
CREATE TABLE IF NOT EXISTS shark (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  type VARCHAR(100),
  max_length INT,
  ocean VARCHAR(100),
  top_speed INT,
  attacks_per_year INT);

-- +migrate Down
DROP TABLE shark;
