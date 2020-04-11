package consts

import "regexp"

// Constants is the structure for all server constants
type Constants struct {
	EmailRegex *regexp.Regexp
}

// MakeConstants makes a new constants instance
func MakeConstants() Constants {
	return Constants{
		EmailRegex: regexp.MustCompile(
			"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
		),
	}
}
