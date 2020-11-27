-- name: GetUserByEmail :one
SELECT Users.User_ID,
  Local_Users.Email,
  Profiles.Username,
  Profiles.Avatar,
  Profiles.Status,
  Local_Users.Password
FROM Local_Users
  INNER JOIN Users ON (Local_Users.User_ID = Users.User_ID)
  INNER JOIN Profiles ON (Local_Users.User_ID = Profiles.User_ID)
WHERE Local_Users.Email = $1;

-- name: AddUser :exec
INSERT INTO Users (User_ID)
VALUES ($1);

-- name: AddProfile :exec
INSERT INTO Profiles(User_ID, Username, Avatar, Status)
VALUES ($1, $2, $3, $4);

-- name: AddLocalUser :exec
INSERT INTO Local_Users (User_ID, Email, Password)
VALUES ($1, $2, $3);

-- name: AddForeignUser :one
INSERT INTO Foreign_Users (User_ID, Home_Server, Local_User_ID)
VALUES ($1, $2, $3) ON CONFLICT (Local_User_ID) DO
UPDATE
SET Local_User_ID = Foreign_Users.Local_User_ID RETURNING Local_User_ID;

-- name: GetUser :one
SELECT Users.User_ID,
  Profiles.Username,
  Profiles.Avatar,
  Profiles.Status
FROM Users
  INNER JOIN Profiles ON (Users.User_ID = Profiles.User_ID)
WHERE Users.User_ID = $1;

-- name: GetLocalUserID :one
SELECT Local_User_ID
FROM Foreign_Users
WHERE User_ID = $1
  AND Home_Server = $2;

-- name: EmailExists :one
SELECT COUNT(*)
FROM Local_Users
WHERE Email = $1;

-- name: UpdateUsername :exec
UPDATE Profiles
SET Username = $1
WHERE User_ID = $2;

-- name: UpdateAvatar :exec
UPDATE Profiles
SET Avatar = $1
WHERE User_ID = $2;

-- name: GetAvatar :one
SELECT Avatar
FROM Profiles
WHERE User_ID = $1;

-- name: SetStatus :exec
UPDATE Profiles
SET Status = $1
WHERE User_ID = $2;

-- name: GetUserMetadata :one
SELECT Metadata
FROM User_Metadata
WHERE User_ID = $1
  AND App_ID = $2;

-- name: AddToGuildList :exec
INSERT INTO Guild_List (User_ID, Guild_ID, Home_Server, Position)
VALUES($1, $2, $3, $4);

-- name: GetGuildListPosition :one
SELECT Position
FROM Guild_List
WHERE User_ID = $1
  AND Guild_ID = $2
  AND Home_Server = $3;

-- name: MoveGuild :exec
UPDATE Guild_List
SET Position = $1
WHERE User_ID = $1
  AND Guild_ID = $2
  AND Home_Server = $3;

-- name: RemoveGuildFromList :exec
DELETE FROM Guild_List
WHERE User_ID = $1
  AND Guild_ID = $2
  AND Home_Server = $3;

-- name: GetGuildList :many
SELECT Guild_ID,
  Home_Server
FROM Guild_List
WHERE User_ID = $1
ORDER BY Position;

-- name: GetLastGuildPositionInList :one
SELECT Position
FROM Guild_List
WHERE User_ID = $1
ORDER BY Position
LIMIT 1;

-- name: UserIsLocal :one
SELECT EXISTS(
    SELECT 1
    FROM Local_Users
    WHERE User_ID = $1
  );

-- name: IsIPWhitelisted :one
SELECT EXISTS (
    SELECT 1
    FROM Rate_Limit_Whitelist_IP
    WHERE IP = $1
  );

-- name: IsUserWhitelisted :one
SELECT EXISTS (
    SELECT 1
    FROM Rate_Limit_Whitelist_User
    WHERE User_ID = $1
  );