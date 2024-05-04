CREATE TABLE IF NOT EXISTS public.users (
	id       SERIAL       NOT NULL UNIQUE,
	username VARCHAR(20)  NOT NULL UNIQUE,
	password VARCHAR(512) NOT NULL, -- a hash of the real password in here
		CONSTRAINT pk_users PRIMARY KEY (id)
);
