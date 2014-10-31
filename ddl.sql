DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS questions CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS options CASCADE;
DROP TABLE IF EXISTS users_categories CASCADE;
DROP TABLE IF EXISTS users_options CASCADE;


CREATE TABLE users (
	id			serial PRIMARY KEY,
	email 		varchar(256) NOT NULL,
	encrypted_password	varchar(256),
	created_at 	timestamp DEFAULT current_timestamp,
	updated_at	timestamp,
	provider	varchar(256),
	uid			varchar(256),
	avatar_path	varchar(256) 
);

CREATE TABLE questions (
	id 				serial PRIMARY KEY,
	text			varchar(256),
	tag				varchar(256),
	display_order	integer DEFAULT 0
);

CREATE TABLE categories (
	id		serial PRIMARY KEY,
	text	varchar(256),
	question_id		integer REFERENCES questions(id),
	display_order 	integer DEFAULT 0
);

CREATE TABLE options (
	id 		serial PRIMARY KEY,
	text	varchar(256),
	category_id		integer REFERENCES categories(id),
	display_order	integer DEFAULT 0
);

CREATE TABLE users_categories (
	id 			serial PRIMARY KEY,
	user_id		integer REFERENCES users(id),
	category_id integer REFERENCES categories(id)
);

CREATE TABLE users_options (
	id 			serial PRIMARY KEY,
	user_id		integer REFERENCES users(id),
	option_id	integer REFERENCES options(id)
);