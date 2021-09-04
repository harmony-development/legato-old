package dynamicauth

import authv1 "github.com/harmony-development/legato/gen/auth/v1"

type ChoiceStep struct {
	*BaseStep
	Options []string
}

func NewChoiceStep(choices []string, id string, canGoBack bool) *ChoiceStep {
	return &ChoiceStep{
		&BaseStep{
			id,
			canGoBack,
		},
		choices,
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
