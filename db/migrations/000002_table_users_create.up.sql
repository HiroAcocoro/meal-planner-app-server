CREATE TABLE IF NOT EXISTS users (
    id          VARCHAR(36)        NOT NULL UNIQUE,
    email       VARCHAR(50)    NOT NULL,
    password    VARCHAR(100)    NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;