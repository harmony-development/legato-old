-- name: ResolveGuildID :one
SELECT Guild_ID FROM Invites
    WHERE Invite_ID = $1;

-- name: IncrementInvite :exec
UPDATE Invites
    SET Uses=Uses + 1
    WHERE Invite_ID = $1;

-- name: DeleteInvite :execrows
DELETE FROM Invites
    WHERE Invite_ID = $1;

-- name: CreateGuildInvite :one
INSERT INTO Invites (
    Name, Possible_Uses, Guild_ID
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: OpenInvites :many
SELECT * FROM Invites
    WHERE Guild_ID = $1
    AND ( Uses < Possible_Uses OR Possible_Uses = -1);