-- name: CreateEmotePack :exec
INSERT INTO Emote_Packs (Pack_ID, Pack_Name, User_ID)
VALUES ($1, $2, $3);

-- name: DeleteEmotePack :exec
DELETE FROM Emote_Packs
WHERE Pack_ID = $1
	AND User_ID = $2;

-- name: GetEmotePacks :many
SELECT Emote_Packs.Pack_ID,
	Emote_Packs.User_ID,
	Emote_Packs.Pack_Name
FROM Emote_Packs
	INNER JOIN Acquired_Emote_Packs ON Acquired_Emote_Packs.Pack_ID = Emote_Packs.Pack_ID
WHERE Acquired_Emote_Packs.User_ID = $1;

-- name: GetEmotePackEmotes :many
SELECT Image_ID,
	Emote_Name
FROM Emote_Pack_Emotes
WHERE Pack_ID = $1;

-- name: AddEmoteToPack :exec
INSERT INTO Emote_Pack_Emotes (Pack_ID, Image_ID, Emote_Name)
VALUES ($1, $2, $3);

-- name: DeleteEmoteFromPack :exec
DELETE FROM Emote_Pack_Emotes
WHERE Pack_ID = $1
	AND Image_ID = $2;

-- name: AcquireEmotePack :exec
INSERT INTO Acquired_Emote_Packs (Pack_ID, User_ID)
VALUES ($1, $2);

-- name: DequipEmotePack :exec
DELETE FROM Acquired_Emote_Packs
WHERE Pack_ID = $1
	AND User_ID = $2;

-- name: GetPackOwner :one
SELECT Pack_ID
FROM Emote_Packs
WHERE Pack_ID = $1;