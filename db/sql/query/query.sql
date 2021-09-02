-- name: GetSession :one
SELECT UserID FROM AuthSessions WHERE SessionID = $1;

-- name: SetSession :exec
INSERT INTO AuthSessions(UserID, SessionID) VALUES($1, $2);

-- name: DeleteSession :exec
DELETE FROM AuthSessions WHERE SessionID = $1;
