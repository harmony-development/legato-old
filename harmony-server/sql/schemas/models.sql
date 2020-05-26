CREATE TABLE IF NOT EXISTS Users
(
    User_ID  BIGSERIAL   NOT NULL,
    Username TEXT UNIQUE NOT NULL,
    Avatar   TEXT,
    PRIMARY KEY (User_ID)
);
CREATE TABLE IF NOT EXISTS Local_Users
(
    User_ID   BIGSERIAL   NOT NULL,
    Email     TEXT UNIQUE NOT NULL,
    Password  BYTEA       NOT NULL,
    Instances JSONB[],
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Foreign_Users
(
    User_ID       BIGSERIAL        NOT NULL,
    Home_Server   TEXT             NOT NULL,
    Local_User_ID BIGSERIAL UNIQUE NOT NULL,
    FOREIGN KEY (Local_User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Sessions
(
    Session    TEXT PRIMARY KEY NOT NULL,
    User_ID    BIGSERIAL        NOT NULL,
    Expiration BIGINT           NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES Users (User_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Guilds
(
    Guild_ID    BIGSERIAL PRIMARY KEY NOT NULL,
    Owner_ID    BIGSERIAL             NOT NULL,
    Guild_Name  TEXT                  NOT NULL,
    Picture_URL TEXT                  NOT NULL
);
CREATE TABLE IF NOT EXISTS Guild_Members
(
    User_ID  BIGSERIAL NOT NULL,
    Guild_ID BIGSERIAL NOT NULL,
    UNIQUE (User_ID, Guild_ID),
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Invites
(
    Invite_ID     BIGSERIAL PRIMARY KEY UNIQUE,
    Name          TEXT      NOT NULL,
    Uses          INTEGER   NOT NULL DEFAULT 0,
    Possible_Uses INTEGER            DEFAULT -1,
    Guild_ID      BIGSERIAL NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Channels
(
    Channel_ID   BIGSERIAL PRIMARY KEY UNIQUE,
    Guild_ID     BIGSERIAL,
    Channel_Name TEXT NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Messages
(
    Message_ID BIGSERIAL PRIMARY KEY,
    Guild_ID   BIGSERIAL NOT NULL,
    Channel_ID BIGSERIAL NOT NULL,
    User_ID    BIGSERIAL NOT NULL,
    Created_At TIMESTAMP NOT NULL,
    Edited_At  TIMESTAMP,
    Content    TEXT      NOT NULL,
    Embeds     jsonb[],
    Actions    jsonb[],
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE,
    FOREIGN KEY (Channel_ID) REFERENCES Channels (Channel_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Attachments
(
    Message_ID BIGSERIAL NOT NULL,
    Attachment TEXT      NOT NULL,
    FOREIGN KEY (Message_ID) REFERENCES Messages (Message_ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS Files
(
    Hash    BYTEA PRIMARY KEY NOT NULL,
    File_ID TEXT              NOT NULL
)