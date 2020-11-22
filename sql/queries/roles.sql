-- name: GetGuildRoles :one
SELECT Roles
    FROM Guilds
    WHERE Guild_ID = $1;

-- name: SetGuildRoles :exec
UPDATE Guilds
    SET Roles = $1
    WHERE Guild_ID = $2;

-- name: GetGuildPerms :one
SELECT Permissions
    FROM Guilds
    WHERE Guild_ID = $1;

-- name: SetGuildPerms :exec
UPDATE Guilds
    SET Permissions = $1
    WHERE Guild_ID = $2;
