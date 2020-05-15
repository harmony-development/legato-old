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