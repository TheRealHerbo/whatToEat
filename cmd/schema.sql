CREATE TABLE IF NOT EXISTS recipie (
  id           INTEGER PRIMARY KEY,
  name         text    NOT NULL,
  ingredients  text,
  directions   text
);