-- name: AddNonce :exec
INSERT INTO Federation_Nonces (Nonce, User_ID, Home_Server)
VALUES ($1, $2, $3);

-- name: GetNonceInfo :one
SELECT User_ID,
  Home_Server
FROM Federation_Nonces
WHERE Nonce = $1;