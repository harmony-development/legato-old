-- SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
--
-- SPDX-License-Identifier: AGPL-3.0-or-later

-- name: GetUserByEmail :one
SELECT
    UserID,
    Passwd
FROM Users WHERE Email=$1 LIMIT 1;
