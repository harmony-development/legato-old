-- name: SessionToUserID :one
SELECT User_ID
FROM Sessions
WHERE Session = $1;

-- name: AddSession :exec
INSERT INTO Sessions (User_ID, Session, Expiration)
VALUES ($1, $2, $3);

-- name: SetExpiration :exec
UPDATE Sessions
SET Expiration = $1
WHERE User_ID = $2;

-- name: AddTimeToSession :exec
UPDATE Sessions
    SET Expiration = (select extract(epoch from now()) + 172800)
    WHERE Session = $1;

-- name: ExpireSessions :exec
DELETE FROM Sessions
WHERE Expiration <= $1;