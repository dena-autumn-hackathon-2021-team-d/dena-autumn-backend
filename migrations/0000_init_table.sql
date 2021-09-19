-- +migrate Up
CREATE TABLE IF NOT EXISTS groups (
  id VARCHAR(50) PRIMARY KEY,
  name VARCHAR(32),
  created_at VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS questions (
  id VARCHAR(50),
  group_id VARCHAR(50),
  contents VARCHAR(256),
  username VARCHAR(32),
  created_at VARCHAR(20),
  PRIMARY KEY (id, group_id),
  FOREIGN KEY (group_id) REFERENCES groups (id)
);

CREATE TABLE IF NOT EXISTS answers (
  id VARCHAR(50),
  group_id VARCHAR(50),
  question_id INT,
  contents VARCHAR(256),
  username VARCHAR(32),
  created_at VARCHAR(20),
  PRIMARY KEY (id, group_id, question_id),
  FOREIGN KEY (group_id) REFERENCES groups (id),
  FOREIGN KEY (question_id) REFERENCES questions (id)
);

CREATE TABLE IF NOT EXISTS comments (
  id VARCHAR(50),
  group_id VARCHAR(50),
  question_id INT,
  answer_id INT,
  contents VARCHAR(256),
  username VARCHAR(32),
  created_at VARCHAR(20),
  PRIMARY KEY (id, group_id, question_id, answer_id),
  FOREIGN KEY (group_id) REFERENCES groups (id),
  FOREIGN KEY (question_id) REFERENCES questions (id),
  FOREIGN KEY (answer_id) REFERENCES answers (id)
);

-- +migrate Down
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS groups;
