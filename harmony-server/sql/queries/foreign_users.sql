-- name: AddForeignUser :exec
INSERT INTO Foreign_Users
    (User_ID, Home_Server, Username, Avatar)
VALUES ($1, $2, $3, $4);

-- name: GetForeignUser :one
SELECT User_ID, Email, Username, Avatar, Password
FROM Users
WHERE Email = $1;