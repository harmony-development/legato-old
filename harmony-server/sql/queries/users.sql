-- name: GetUser :one
SELECT User_ID, Email, Username, Avatar, Password
FROM Users
WHERE Email = $1;

-- name: AddUser :exec
INSERT INTO Users (User_ID, Email, Username, Avatar, Password)
VALUES ($1, $2, $3, $4, $5);

-- name: GetUserByID :one
SELECT User_ID, Email, Username, Avatar, Password
FROM Users
WHERE User_ID = $1;

-- name: EmailExists :one
SELECT User_ID
FROM Users
WHERE Email = $1;

-- name: UpdateUsername :exec
UPDATE Users
SET Username=$1
WHERE User_ID = $2;

-- name: GetAvatar :one
SELECT Avatar
FROM Users
WHERE User_ID = $1;

-- name: UpdateAvatar :exec
UPDATE Users
SET Avatar=$1
WHERE User_ID = $2;