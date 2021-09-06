// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package dynamicauth

import authv1 "github.com/harmony-development/legato/gen/auth/v1"

type FormStep struct {
	*BaseStep
	Fields []FormField
}

type FormField struct {
	Name      string
	FieldType string
}

func NewFormStep(fields []FormField, id string, canGoBack bool) *FormStep {
	return &FormStep{
		&BaseStep{
			StepTypeForm,
			id,
			canGoBack,
		},
		fields,
	}
}

func (s *FormStep) ToProtoV1() *authv1.AuthStep {
	fields := make([]*authv1.AuthStep_Form_FormField, len(s.Fields))

	for i, f := range s.Fields {
		fields[i] = &authv1.AuthStep_Form_FormField{
			Name: f.Name,
			Type: f.FieldType,
		}
	}

	return &authv1.AuthStep{
		Step: &authv1.AuthStep_Form_{
			Form: &authv1.AuthStep_Form{
				Title:  s.BaseStep.id,
				Fields: fields,
			},
		},
	}
}
