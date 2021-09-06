-- SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
--
-- SPDX-License-Identifier: AGPL-3.0-or-later

-- name: GetSession :one
SELECT UserID FROM AuthSessions WHERE SessionID = $1;

-- name: AddSession :exec
INSERT INTO AuthSessions(UserID, SessionID) VALUES($1, $2);

-- name: DeleteSession :exec
DELETE FROM AuthSessions WHERE SessionID = $1;
