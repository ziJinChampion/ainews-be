CREATE TABLE IF NOT EXISTS tags(
   id serial PRIMARY KEY,
   tag_name VARCHAR (50) UNIQUE NOT NULL,
   tag_description text,

   created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS articles(
   id serial PRIMARY KEY,
   tag_id int NOT NULL ,
   content text,
   author_id int NOT NULL,
   title text NOT NULL,
   is_hidden boolean NOT NULL default false,

   created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP
);


