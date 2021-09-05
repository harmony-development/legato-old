// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package dynamicauth

import authv1 "github.com/harmony-development/legato/gen/auth/v1"

type Step interface {
	ID() string
	CanGoBack() bool
	ToProtoV1() *authv1.AuthStep
}

type BaseStep struct {
	id        string
	canGoBack bool
}

func (s *BaseStep) ID() string {
	return s.id
}

func (s *BaseStep) CanGoBack() bool {
	return s.canGoBack
}
