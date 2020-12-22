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

-- name: PermissionExistsWithoutChannel :one
SELECT EXISTS(SELECT 1 FROM PERMISSIONS WHERE Guild_ID = $1 AND Channel_ID IS NULL AND Role_ID = $2);

-- name: PermissionExistsWithoutChannelWithoutRole :one
SELECT EXISTS(SELECT 1 FROM PERMISSIONS WHERE Guild_ID = $1 AND Channel_ID IS NULL AND Role_ID IS NULL);

-- name: PermissionsExists :one
SELECT EXISTS(SELECT 1 FROM PERMISSIONS WHERE Guild_ID = $1 AND Channel_ID = $2 AND Role_ID = $3);

-- name: PermissionsExistsWithoutRole :one
SELECT EXISTS(SELECT 1 FROM PERMISSIONS WHERE Guild_ID = $1 AND Channel_ID = $2 AND Role_ID IS NULL);

-- name: UpdatePermissions :exec
UPDATE Permissions
SET Nodes = $4
WHERE Guild_ID = $1
    AND Channel_ID = $2
    AND Role_ID = $3;

-- name: UpdatePermissionsWithoutRole :exec
UPDATE Permissions
SET Nodes = $3
WHERE Guild_ID = $1
    AND Channel_ID = $2
    AND Role_ID IS NULL;

-- name: UpdatePermissionsWithoutChannel :exec
UPDATE Permissions
SET Nodes = $3
WHERE Guild_ID = $1
    AND Channel_ID IS NULL
    AND Role_ID = $2;

-- name: UpdatePermissionsWithoutChannelWithoutRole :exec
UPDATE Permissions
SET Nodes = $2
WHERE Guild_ID = $1
    AND Channel_ID IS NULL
    AND Role_ID IS NULL;

-- name: SetPermissions :exec
INSERT INTO Permissions
    (Guild_ID, Channel_ID, Role_ID, Nodes)
    VALUES ($1, $2, $3, $4);

-- name: GetPermissions :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Channel_ID = $2
    AND Role_ID = $3;

-- name: GetPermissionsWithoutRole :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Channel_ID = $2
    AND Role_ID IS NULL;

-- name: GetPermissionsWithoutChannel :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Channel_ID IS NULL
    AND Role_ID = $2;

-- name: GetPermissionsWithoutChannelWithoutRole :one
SELECT Nodes
FROM Permissions
WHERE Guild_ID = $1
    AND Channel_ID IS NULL
    AND Role_ID IS NULL;

-- name: SetRoleName :exec
UPDATE Roles
    SET Name = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3;

-- name: SetRoleColor :exec
UPDATE Roles
    SET Color = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3;

-- name: SetRoleHoist :exec
UPDATE Roles
    SET Hoist = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3;

-- name: SetRolePingable :exec
UPDATE Roles
    SET Pingable = $1
    WHERE Guild_ID = $2
      AND Role_ID = $3;
