CREATE TABLE IF NOT EXISTS guilds(guildid TEXT PRIMARY KEY UNIQUE, guildname TEXT, picture TEXT);
CREATE TABLE IF NOT EXISTS guildmembers(userid TEXT, guildid TEXT UNIQUE REFERENCES guilds(guildid));
CREATE TABLE IF NOT EXISTS users(id TEXT PRIMARY KEY UNIQUE, email TEXT UNIQUE, username TEXT, avatar TEXT, password TEXT);
CREATE TABLE IF NOT EXISTS invites(inviteid TEXT PRIMARY KEY UNIQUE, guildid TEXT REFERENCES guilds(guildid));
INSERT INTO guilds(guildid, guildname, picture) VALUES("harmony-devs", "Harmony Development", "") ON CONFLICT DO NOTHING;
INSERT INTO invites(inviteid, guildid) VALUES("join-harmony-dev", "harmony-dev") ON CONFLICT DO NOTHING;
INSERT INTO users(id, email, username, avatar, password) VALUES("82ee9c8dc9e165205548b7c3833e7372", "developer@harmonyapp.io", "developer", "", "$2a$10$WHuq8sNHk0ks0JwlpkV36eNmpEvD7r9pqI/F7kB0q0yAUpENzmtne") ON CONFLICT DO NOTHING;
INSERT INTO users(id, email, username, avatar, password) VALUES("dadcd6bf8c0338cbfc9aa9c369ea93cc", "developer2@harmonyapp.io", "developer2", "", "$2a$10$yTHVSHmbAAgcIysrJZg/cesPg7o9qpoTGxFgeM/7pQIgOLFjJZPLW") ON CONFLICT DO NOTHING;
