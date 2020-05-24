-- name: AddFileHash :exec
INSERT INTO Files(Hash, File_ID)
VALUES ($1, $2);

-- name: GetFileByHash :one
SELECT File_ID
FROM Files
WHERE Hash = $1;