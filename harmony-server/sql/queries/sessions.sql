-- name: SessionToUserID :one
SELECT User_ID
FROM Sessions
WHERE Session = $1;

-- name: AddSession :exec
INSERT INTO Sessions (User_ID,
                      Session)
VALUES ($1, $2);