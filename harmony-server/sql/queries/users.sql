-- name: GetUserByEmail :one
SELECT Users.User_ID, Local_Users.Email, Username, Avatar, Local_Users.Password
FROM Local_Users
         INNER JOIN Users
                    ON (Local_Users.User_ID = Users.User_ID)
WHERE Local_Users.Email = $1;

-- name: AddUser :exec
INSERT INTO Users (User_ID, Home_Server, Username, Avatar)
VALUES ($1, $2, $3, $4);

-- name: AddLocalUser :exec
INSERT INTO Local_Users (Email, Password, Instances)
VALUES ($1, $2, $3);

-- name: GetUser :one
SELECT User_ID, Username, Avatar
FROM Users
WHERE User_ID = $1
  AND Home_Server = $2;

-- name: EmailExists :one
SELECT User_ID
FROM Local_Users
WHERE Email = $1;

-- name: UpdateUsername :exec
UPDATE Users
SET Username=$1
WHERE User_ID = $2
  AND Home_Server = $3;

-- name: UpdateAvatar :exec
UPDATE Users
SET Avatar=$1
WHERE User_ID = $2
  AND Home_Server = $3;

-- name: GetAvatar :one
SELECT Avatar
FROM Users
WHERE User_ID = $1
  AND Home_Server = $2;