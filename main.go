// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"log"
	"os"

	"github.com/harmony-development/legato/logger"
	"github.com/harmony-development/legato/server"
)

func main() {
	l := logger.New(os.Stdin)

	server, err := server.New(l)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Listen())
}
