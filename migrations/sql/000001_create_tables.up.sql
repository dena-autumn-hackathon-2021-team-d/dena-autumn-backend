BEGIN;

CREATE TABLE IF NOT EXISTS questions (
  id INT NOT NULL PRIMARY KEY,
  question VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS answers (
  id INT NOT NULL PRIMARY KEY,
  question_id INT NOT NULL,
  answer VARCHAR(256),
  FOREIGN KEY fk_answers_questions (question_id)
    REFERENCES questions (id)
);

COMMIT;
