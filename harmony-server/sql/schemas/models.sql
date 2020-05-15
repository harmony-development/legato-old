CREATE TABLE Sessions (
    User_ID TEXT NOT NULL,
    Session TEXT PRIMARY KEY NOT NULL
);
CREATE TABLE Guilds (
    Guild_ID BIGSERIAL PRIMARY KEY NOT NULL,
    Owner_ID TEXT NOT NULL,
    Guild_Name TEXT NOT NULL,
    Picture_URL TEXT NOT NULL
);
CREATE TABLE Guild_Members (
    User_ID TEXT NOT NULL,
    Guild_ID BIGSERIAL NOT NULL,
    UNIQUE(User_ID, Guild_ID),
    FOREIGN KEY (Guild_ID) REFERENCES Guilds(Guild_ID)
);
CREATE TABLE Invites (
    Invite_ID BIGSERIAL PRIMARY KEY UNIQUE,
    Uses INTEGER NOT NULL DEFAULT 0,
    Guild_ID BIGSERIAL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds(Guild_ID)
);
CREATE TABLE Channels (
    Channel_ID BIGSERIAL PRIMARY KEY UNIQUE,
    Guild_ID BIGSERIAL,
    Channel_Name TEXT NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds(Guild_ID)
);
CREATE TABLE Messages (
    Message_ID BIGSERIAL PRIMARY KEY,
    Guild_ID BIGSERIAL,
    Channel_ID BIGSERIAL,
    User_ID TEXT NOT NULL,
    Created_At TIMESTAMP NOT NULL,
    Edited_At TIMESTAMP,
    Content TEXT NOT NULL,
    FOREIGN KEY (Guild_ID) REFERENCES Guilds(Guild_ID),
    FOREIGN KEY (Channel_ID) REFERENCES Channels(Channel_ID)
);
CREATE TABLE Attachments (
    Message_ID TEXT NOT NULL,
    Attachment_URL TEXT NOT NULL,
    FOREIGN KEY (Message_ID) REFERENCES  Messages(Message_ID)
);