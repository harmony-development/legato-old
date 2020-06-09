-- name: SessionToUserID :one
SELECT User_ID
FROM Sessions
WHERE Session = $1;

-- name: AddSession :exec
INSERT INTO Sessions
(User_ID,
 Session,
 Expiration)
VALUES ($1, $2, $3);

-- name: ExpireSessions :exec
DELETE
FROM Sessions
WHERE Expiration <= $1;