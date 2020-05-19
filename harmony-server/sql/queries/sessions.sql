-- name: SessionToUserID :one
SELECT User_ID, Home_Server
FROM Sessions
WHERE Session = $1;

-- name: AddSession :exec
INSERT INTO Sessions
(User_ID,
 Home_Server,
 Session,
 Expiration)
VALUES ($1, $2, $3, $4);

-- name: ExpireSessions :exec
DELETE
FROM Sessions
WHERE Expiration <= $1;