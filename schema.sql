CREATE TABLE IF NOT EXISTS servers
(
    id TEXT NOT NULL PRIMARY KEY UNIQUE,
    icon TEXT
);

CREATE TABLE IF NOT EXISTS server_messages
(
    server_id TEXT REFERENCES servers,
    message TEXT,
    message_id TEXT NOT NULL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS message_attachments
(
    message_id TEXT UNIQUE REFERENCES server_messages,
    attachment TEXT
);

CREATE TABLE IF NOT EXISTS users
(
    id TEXT NOT NULL PRIMARY KEY UNIQUE,
    email TEXT NOT NULL UNIQUE ,
    password TEXT NOT NULL ,
    username TEXT,
    avatar TEXT
);

CREATE TABLE IF NOT EXISTS user_servers
(
    user_id TEXT UNIQUE REFERENCES users,
    server_id TEXT REFERENCES servers
);

