-- name: CreateGuild :one
INSERT INTO Guilds (Owner_ID, Guild_Name, Picture_URL)
VALUES ($1, $2, $3)
RETURNING *;

-- name: AddUserToGuild :exec
INSERT INTO Guild_Members (User_ID, Guild_ID)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: CreateChannel :one
INSERT INTO Channels (Guild_ID, Channel_Name)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteGuild :exec
DELETE
FROM Guilds
WHERE Guild_ID = $1;

-- name: GetGuildOwner :one
SELECT Owner_ID
from GUILDS
WHERE Guild_ID = $1;

-- name: CreateGuildInvite :one
INSERT INTO Invites (Name, Possible_Uses, Guild_ID)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteChannel :exec
DELETE
FROM Channels
WHERE Guild_ID = $1
  AND Channel_ID = $2;

-- name: AddMessage :one
INSERT INTO Messages(Message_ID, Guild_ID, Channel_ID, User_ID, Created_At, Edited_At, Content)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: AddAttachment :exec
INSERT INTO Attachments(Message_ID, Attachment_URL)
VALUES ($1, $2);

-- name: DeleteMessage :execrows
DELETE
FROM Messages
WHERE Message_ID = $1
  AND Channel_ID = $2
  AND Guild_ID = $3;

-- name: GetMessageAuthor :one
SELECT User_ID
FROM Messages
WHERE Message_ID = $1;

-- name: ResolveGuildID :one
SELECT Guild_ID
FROM Invites
WHERE Invite_ID = $1;

-- name: IncrementInvite :exec
UPDATE Invites
SET Uses=Uses + 1
WHERe Invite_ID = $1;

-- name: DeleteInvite :execrows
DELETE
FROM Invites
WHERE Invite_ID = $1;

-- name: SessionToUserID :one
SELECT User_ID
FROM Sessions
WHERE Session = $1;

-- name: UserInGuild :one
SELECT User_ID
FROM Guild_Members
WHERE User_ID = $1
  AND Guild_ID = $2;

-- name: GetAttachments :many
SELECT Attachment_URL
FROM Attachments
WHERE Message_ID = $1;

-- name: GetMessageDate :one
SELECT Created_At
FROM Messages
WHERE Message_ID = $1;

-- name: GetMessages :many
SELECT Message_ID, User_ID, Content, Created_At
FROM Messages
WHERE Guild_ID = @GuildID
  AND Channel_ID = @ChannelID
  AND Created_At < @Before
ORDER BY Created_At DESC
LIMIT @Max;