// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package errwrap

import "fmt"

func Wrap(err error, wrap string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", wrap, err)
}

func Wrapf(err error, wrap string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(wrap, args...))
}
