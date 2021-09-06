-- SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
--
-- SPDX-License-Identifier: AGPL-3.0-or-later

CREATE TABLE AuthSessions (
    SessionID TEXT PRIMARY KEY,
    UserID BIGINT NOT NULL
);
