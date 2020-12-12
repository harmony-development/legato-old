-- name: AddFileMetadata :exec
INSERT INTO Files (File_ID, Content_Type, Name, Size)
VALUES ($1, $2, $3, $4);

-- name: GetFileMetadata :one
SELECT Content_Type,
	Name,
	Size
FROM Files
WHERE File_ID = $1;

-- name: GetFileIDByHash :one
SELECT File_ID
FROM Hashes
WHERE Hash = $1;

-- name: AddHash :exec
INSERT INTO Hashes (Hash, File_ID)
VALUES ($1, $2);

-- name: DeleteFileMetadata :exec
DELETE FROM Files
WHERE File_ID = $1;