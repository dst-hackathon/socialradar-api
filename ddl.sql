CREATE TABLE USER (
	id			integer PRIMARY KEY,
	email 		varchar(256) NOT NULL,
	encrypted_password	varchar(256),
	created_at 	timestamp DEFAULT current_timestamp,
	updated_at	timestamp,
	provider	varchar(256),
	uid			varchar(256),
	avatar_path	varchar(256) 
)

CREATE TABLE QUESTION (
	id 		integer PRIMARY KEY,
	text	varchar(256),
	display_order	integer DEFAULT 0
)

CREATE TABLE CATEGORY (
	id		integer PRIMARY KEY,
	text	varchar(256),
	question_id		integer REFERENCES QUESTION(id),
	display_order 	integer DEFAULT 0
)

CREATE TABLE OPTION (
	id 		integer PRIMARY KEY,
	text	varchar(256),
	category_id		integer REFERENCES CATEGORY(id),
	display_order	integer DEFAULT 0
)

CREATE TABLE USER_CATEGORY (
	id 			integer PRIMARY KEY,
	user_id		integer REFERENCES USER(id),
	category_id integer REFERENCES CATEGORY(id)
)

CREATE TABLE USER_OPTION (
	id 			integer PRIMARY KEY,
	user_id		integer REFERENCES USER(id),
	option_id	integer REFERENCES OPTION(id)
)