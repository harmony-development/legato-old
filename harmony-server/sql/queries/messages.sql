-- name: AddMessage :one
INSERT INTO Messages (Guild_ID, Channel_ID, User_ID, Content, Embeds, Actions, Created_At)
VALUES ($1, $2, $3, $4, $5, $6, NOW())
RETURNING *;

-- name: AddAttachment :exec
INSERT INTO Attachments(Message_ID, Attachment)
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

-- name: GetMessage :one
SELECT *
FROM Messages
WHERE Message_ID = $1;

-- name: GetAttachments :many
SELECT Attachment
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