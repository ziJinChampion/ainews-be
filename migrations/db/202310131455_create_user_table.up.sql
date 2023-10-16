CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   user_name VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (50) ,
   mobile VARCHAR (50) ,
   user_role VARCHAR(30) NOT NULL default 'NORMAL_USER',
   grade int NOT NULL default 1,
   created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP 
);