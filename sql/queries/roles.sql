-- name: CreateRole :one
INSERT INTO Roles (
        Guild_ID,
        Role_ID,
        Name,
        Color,
        Hoist,
        Pingable,
        Position
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7
    ) RETURNING *;

-- name: DeleteRole :exec
DELETE FROM Roles
WHERE Guild_ID = $1
    AND Role_ID = $2;

-- name: GetRolesForGuild :many
SELECT *
FROM Roles
WHERE Guild_ID = $1
ORDER BY Position;

-- name: MoveRole :exec
UPDATE Roles
SET Position = $1
WHERE Role_ID = $2
    AND Guild_ID = $3;

-- name: GetRolePosition :one
SELECT Position
FROM Roles
WHERE Role_ID = $1
    AND Guild_ID = $2;

-- name: RolesForUser :many
SELECT Role_ID
FROM Roles_Members
WHERE Guild_ID = $1
    AND Member_ID = $2;

-- name: AddUserToRole :exec
INSERT INTO Roles_Members (Guild_ID, Role_ID, Member_ID)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;

-- name: RemoveUserFromRole :exec
DELETE FROM Roles_Members
WHERE Guild_ID = $1
    AND Role_ID = $2
    AND Member_ID = $3;

-- name: SetPermissions :exec
INSERT INTO Permissions (Guild_ID, Channel_ID, Role_ID, Nodes)
VALUES ($1, $2, $3, $4) ON CONFLICT (Guild_ID, Channel_ID, Role_ID) DO
UPDATE
SET Nodes = EXCLUDED.Nodes;

-- name: GetPermissions :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Channel_ID = $2
    AND Role_ID = $3;

-- name: GetPermissionsWithoutChannel :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Role_ID = $2;