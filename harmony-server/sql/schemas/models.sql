CREATE TABLE Sessions
(
    User_ID BIGSERIAL        NOT NULL,
    Session TEXT PRIMARY KEY NOT NULL
);
CREATE TABLE Guilds
(
    Guild_ID    BIGSERIAL PRIMARY KEY NOT NULL,
    Owner_ID    BIGSERIAL             NOT NULL,
    Guild_Name  TEXT                  NOT NULL,
    Picture_URL TEXT                  NOT NULL
);
CREATE TABLE Guild_Members
(
    User_ID  BIGSERIAL NOT NULL,
    Guild_ID BIGSERIAL NOT NULL,
    UNIQUE (User_ID, Guild_ID),
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);
CREATE TABLE Invites
(
    Invite_ID     BIGSERIAL PRIMARY KEY UNIQUE,
    Name          TEXT      NOT NULL,
    Uses          INTEGER   NOT NULL DEFAULT 0,
    Possible_Uses INTEGER            DEFAULT -1,
    Guild_ID      BIGSERIAL NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);
CREATE TABLE Channels
(
    Channel_ID   BIGSERIAL PRIMARY KEY UNIQUE,
    Guild_ID     BIGSERIAL,
    Channel_Name TEXT NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds (Guild_ID) ON DELETE CASCADE
);
CREATE TABLE Messages
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
CREATE TABLE Attachments
(
    Message_ID     BIGSERIAL NOT NULL,
    Attachment_URL TEXT      NOT NULL,
    FOREIGN KEY (Message_ID) REFERENCES Messages (Message_ID) ON DELETE CASCADE
);
