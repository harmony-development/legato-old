-- name: GetUser :one
SELECT User_ID, Email, Username, Avatar, Password
FROM Users
WHERE Email = $1;

-- name: AddUser :exec
INSERT INTO Users (User_ID, Email, Username, Avatar, Password)
VALUES ($1, $2, $3, $4, $5);

-- name: EmailExists :one
SELECT User_ID
FROM Users
WHERE Email = $1;