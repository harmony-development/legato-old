// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package dynamicauth

import authv1 "github.com/harmony-development/legato/gen/auth/v1"

type ChoiceStep struct {
	*BaseStep
	Options   []string
	optionMap map[string]struct{}
}

func NewChoiceStep(options []string, id string, canGoBack bool) *ChoiceStep {
	optionMap := map[string]struct{}{}
	for _, o := range options {
		optionMap[o] = struct{}{}
	}

	return &ChoiceStep{
		&BaseStep{
			StepTypeChoice,
			id,
			canGoBack,
		},
		options,
		optionMap,
	}
}

func (c *ChoiceStep) ToProtoV1() *authv1.AuthStep {
	return &authv1.AuthStep{
		CanGoBack: c.canGoBack,
		Step: &authv1.AuthStep_Choice_{
			Choice: &authv1.AuthStep_Choice{
				Title:   c.id,
				Options: c.Options,
			},
		},
	}
}

func (c *ChoiceStep) HasOption(option string) bool {
	_, ok := c.optionMap[option]
	return ok
}
