-- SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
--
-- SPDX-License-Identifier: AGPL-3.0-or-later

CREATE TABLE IF NOT EXISTS AuthSessions (
    SessionID TEXT PRIMARY KEY,
    UserID BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS Users (
    UserID BIGINT PRIMARY KEY NOT NULL,
    Email TEXT NOT NULL UNIQUE,
    -- Password is a keyword so uhhh
    Passwd BYTEA NOT NULL
);
