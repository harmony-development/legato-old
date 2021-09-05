// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"github.com/harmony-development/legato/server"
)

func main() {
	it := server.ProduceServer()
	it.Listen()
}
