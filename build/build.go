// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package build

// GitCommit is injected via ldflags.
// This cannot be made a constant so the global check is invalid here
// nolint
var GitCommit string

// Version is the version of legato running.
const Version = "0.0.1"
