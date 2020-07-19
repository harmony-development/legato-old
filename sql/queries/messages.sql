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
SELECT *
FROM Messages
WHERE Guild_ID = @GuildID
  AND Channel_ID = @ChannelID
  AND Created_At < @Before
ORDER BY Created_At DESC
LIMIT @Max;

-- name: UpdateMessageContent :one
UPDATE Messages
  SET Content = $2,
      Edited_At = NOW()
  WHERE Message_ID = $1
RETURNING Content, Edited_At;

-- name: UpdateMessageEmbeds :one
UPDATE Messages
  SET Embeds = $2,
      Edited_At = NOW()
  WHERE Message_ID = $1
RETURNING Embeds, Edited_At;

-- name: UpdateMessageActions :one
UPDATE Messages
  SET Actions = $2,
      Edited_At = NOW()
  WHERE Message_ID = $1
RETURNING Actions, Edited_At;

-- name: MessageWithIDExists :one
SELECT EXISTS (
  SELECT 1 FROM Messages
          WHERE Guild_ID = $1
            AND Channel_ID = $2
            AND Message_ID = $3
);