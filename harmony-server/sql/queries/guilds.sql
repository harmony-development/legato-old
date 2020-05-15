-- name: CreateGuild :one
INSERT INTO Guilds (
    Owner_ID, Guild_Name, Picture_URL
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: AddUserToGuild :exec
INSERT INTO Guild_Members (
    User_ID, Guild_ID
) VALUES (
    $1, $2
)
ON CONFLICT DO NOTHING;

-- name: CreateChannel :one
INSERT INTO Channels (
    Guild_ID, Channel_Name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteGuild :exec
DELETE FROM Guilds WHERE Guild_ID = $1;

-- name: GetGuildOwner :one
SELECT Owner_ID from GUILDS
    WHERE Guild_ID = $1;

-- name: CreateGuildInvite :one
INSERT INTO Invites (
    Name, Possible_Uses, Guild_ID
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeleteChannel :exec
DELETE FROM Channels
    WHERE Guild_ID = $1
    AND Channel_ID = $2;