// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package persist

import "context"

type Sessions interface {
	Get(ctx context.Context, sessionID string) (userID uint64, err error)
	Add(ctx context.Context, sessionID string, userID uint64) error
}
