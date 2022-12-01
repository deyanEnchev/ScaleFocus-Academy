CREATE TABLE IF NOT EXISTS login_info
(
    id integer NOT NULL,
    username VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    CONSTRAINT login_info_pkey PRIMARY KEY (id)
);

-- ALTER TABLE login_info
-- CHANGE password password VARCHAR(256)
-- CHARACTER SET utf8mb4 
-- COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS user_lists
(
    id integer NOT NULL,
    user_id integer NOT NULL,
    name VARCHAR(256) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES login_info (id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_tasks
(
    id integer NOT NULL,
	list_id integer NOT NULL,
    text VARCHAR(256) NOT NULL,
    completed bool NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (list_id) REFERENCES user_lists (id) 
    ON DELETE CASCADE
);