-- name: CreateGuild :one
INSERT INTO Guilds (
    Guild_ID, Owner_ID, Guild_Name, Picture_URL
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetGuildData :one
SELECT * FROM Guilds
    WHERE Guild_ID = $1;

-- name: DeleteGuild :exec
DELETE FROM Guilds
    WHERE Guild_ID = $1;

-- name: AddUserToGuild :exec
INSERT INTO Guild_Members (
    User_ID, Guild_ID
) VALUES (
    $1, $2
)
ON CONFLICT DO NOTHING;

-- name: RemoveUserFromGuild :exec
DELETE FROM Guild_Members
    WHERE Guild_ID = $1
      AND User_ID = $2;

-- name: CreateChannel :one
INSERT INTO Channels (
    Guild_ID, Channel_Name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateChannelName :exec
UPDATE Channels
      SET Channel_Name = $1
    WHERE Guild_ID = $2
      AND Channel_ID = $3;

-- name: GetChannels :many
SELECT * FROM Channels
    WHERE Guild_ID = $1;

-- name: GetGuildOwner :one
SELECT Owner_ID FROM GUILDS
    WHERE Guild_ID = $1;

-- name: DeleteChannel :exec
DELETE FROM Channels
    WHERE Guild_ID = $1
    AND Channel_ID = $2;

-- name: SetGuildName :exec
UPDATE Guilds
    SET Guild_Name = $1
    WHERE Guild_ID = $2;

-- name: GetGuildPicture :one
SELECT Picture_URL FROM Guilds
    WHERE Guild_ID = $1;

-- name: SetGuildPicture :exec
UPDATE Guilds
    SET Picture_URL = $1
    WHERE Guild_ID = $2;

-- name: GetGuildMembers :many
SELECT User_ID FROM Guild_Members
    WHERE Guild_ID = $1;

-- name: GuildsForUser :many
SELECT Guilds.Guild_ID FROM Guild_Members
    INNER JOIN guilds
    ON Guild_Members.Guild_ID = Guilds.Guild_ID
    WHERE User_ID = $1;

-- name: GuildsForUserWithData :many
SELECT * FROM Guild_Members
    INNER JOIN guilds
    ON Guild_Members.Guild_ID = Guilds.Guild_ID
    WHERE User_ID = $1;

-- name: UserInGuild :one
SELECT User_ID FROM Guild_Members
    WHERE User_ID = $1
    AND Guild_ID = $2;

-- name: GuildWithIDExists :one
SELECT EXISTS (
    SELECT 1 FROM Guilds
             WHERE Guild_ID = $1
);

-- name: NumChannelsWithID :one
SELECT COUNT(*) FROM Channels
    WHERE Guild_ID = $1
      AND Channel_ID = $2;