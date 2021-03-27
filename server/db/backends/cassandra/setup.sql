CREATE KEYSPACE IF NOT EXISTS harmony WITH replication = { 'class': 'SimpleStrategy',
'replication_factor': 1 };

CREATE TABLE IF NOT EXISTS harmony.users (
	UserID BIGINT,
	LocalUserID BIGINT,
	HomeServer TEXT,
	Email TEXT,
	Username TEXT,
	Avatar TEXT,
	PRIMARY KEY (UserID, Username)
);