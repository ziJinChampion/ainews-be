CREATE TABLE IF NOT EXISTS article_tag(
   id serial PRIMARY KEY,
   tag_id int NOT NULL,
   article_id int NOT NULL,

   created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP
);
