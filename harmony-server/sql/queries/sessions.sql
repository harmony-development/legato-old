-- name: SessionToUserID :one
SELECT User_ID FROM Sessions
    WHERE Session = $1;