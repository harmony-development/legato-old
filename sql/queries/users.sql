-- name: GetUserByEmail :one
SELECT Users.User_ID, Local_Users.Email, Profiles.Username, Profiles.Avatar, Profiles.Status, Local_Users.Password
FROM Local_Users
         INNER JOIN Users
                    ON (Local_Users.User_ID = Users.User_ID)
         INNER JOIN Profiles
                    ON (Local_Users.User_ID = Profiles.User_ID)
WHERE Local_Users.Email = $1;

-- name: AddUser :exec
INSERT INTO Users (User_ID)
VALUES ($1);

-- name: AddProfile :exec
INSERT INTO Profiles(User_ID, Username, Avatar, Status)
VALUES ($1, $2, $3, $4);

-- name: AddLocalUser :exec
INSERT INTO Local_Users (User_ID, Email, Password, Instances)
VALUES ($1, $2, $3, $4);

-- name: AddForeignUser :one
INSERT INTO Foreign_Users (User_ID, Home_Server, Local_User_ID)
VALUES ($1, $2, $3)
ON CONFLICT (Local_User_ID) DO UPDATE
    SET Local_User_ID=Foreign_Users.Local_User_ID
RETURNING Local_User_ID;

-- name: GetUser :one
SELECT Users.User_ID, Profiles.Username, Profiles.Avatar, Profiles.Status
FROM Users
         INNER JOIN Profiles ON (Users.User_ID = Profiles.User_ID)
WHERE Users.User_ID = $1;

-- name: GetLocalUserID :one
SELECT Local_User_ID
FROM Foreign_Users
WHERE User_ID = $1
  AND Home_Server = $2;

-- name: EmailExists :one
SELECT COUNT(*) FROM Local_Users
  WHERE Email = $1;

-- name: UpdateUsername :exec
UPDATE Profiles
SET Username=$1
WHERE User_ID = $2;

-- name: UpdateAvatar :exec
UPDATE Profiles
SET Avatar=$1
WHERE User_ID = $2;

-- name: GetAvatar :one
SELECT Avatar
FROM Profiles
WHERE User_ID = $1;

-- name: SetStatus :exec
UPDATE Profiles
SET Status=$1
WHERE User_ID = $2;
