#!/usr/bin/env ruby

require 'pg'
require __dir__+'/migration_utils.rb'

conn = connect
conn.transaction do |con|
    con.exec %{
ALTER TABLE Guilds
    DROP COLUMN Roles;
ALTER TABLE Guilds
    DROP COLUMN Permissions;

CREATE TABLE IF NOT EXISTS Roles (
    Guild_ID BIGSERIAL NOT NULL,
    Role_ID BIGSERIAL NOT NULL,
    Name TEXT NOT NULL,
    Color INTEGER NOT NULL,
    Hoist BOOLEAN NOT NULL,
    Pingable BOOLEAN NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    PRIMARY KEY (Role_ID)
);

CREATE TABLE IF NOT EXISTS Roles_Members (
    Guild_ID BIGSERIAL NOT NULL,
    Role_ID BIGSERIAL NOT NULL,
    Member_ID BIGSERIAL NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    FOREIGN KEY (Role_ID) REFERENCES Roles (Role_ID) ON DELETE CASCADE,
    FOREIGN KEY (Member_ID) REFERENCES Users (User_ID) ON DELETE CASCADE,
    PRIMARY KEY (Guild_ID, Role_ID, Member_ID)
);

CREATE TYPE PermissionsNode AS (
    Node TEXT,
    Allow BOOLEAN
);

CREATE TABLE IF NOT EXISTS Permissions (
    Guild_ID BIGSERIAL NOT NULL,
    Channel_ID BIGINT,
    Role_ID BIGSERIAL NOT NULL,
    Nodes PermissionsNode[] NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    FOREIGN KEY (Role_ID) REFERENCES Roles (Role_ID) ON DELETE CASCADE,
    FOREIGN KEY (Channel_ID) REFERENCES Channels (Channel_ID) ON DELETE CASCADE,
    UNIQUE (Guild_ID, Channel_ID, Role_ID)
);
}
end
