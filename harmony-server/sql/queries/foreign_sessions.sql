-- name: ForeignSessionToUserID :one
SELECT User_ID
FROM Foreign_Sessions
WHERE Session = $1;

-- name: AddForeignSession :exec
INSERT INTO Foreign_Sessions
(User_ID,
 Home_Server,
 Session,
 Expiration)
VALUES ($1, $2, $3, $4);

-- name: ExpireForeignSessions :exec
DELETE
FROM Foreign_Sessions
WHERE Expiration <= $1;