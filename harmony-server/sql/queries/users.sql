-- name: GetUser :one
SELECT User_ID, Email, Username, Avatar, Password
FROM Users
WHERE Email = $1;