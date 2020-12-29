package authsteps

// StepType ...
type StepType string

const (
	// StepChoice ...
	StepChoice StepType = "choice"
	// StepForm ...
	StepForm = "form"
)

// Step is an interface which enables polymorphism
type Step interface {
	ID() string
	StepType() StepType
	SubSteps() []Step
}

// BaseStep is the base implementation for the Step interface
type BaseStep struct {
	id       string
	stepType StepType
	subSteps []Step
}

func (b BaseStep) ID() string {
	return b.id
}

func (b BaseStep) StepType() StepType {
	return b.stepType
}

func (b BaseStep) SubSteps() []Step {
	return b.subSteps
}

// FormField ...
type FormField struct {
	Name      string
	FieldType string
}

// FormStep ...
type FormStep struct {
	BaseStep
	Fields []FormField
}

// ChoiceStep ...
type ChoiceStep struct {
	BaseStep
	Choices []string
}

func NewFormStep(id string, fields []FormField, next []Step) FormStep {
	return FormStep{
		BaseStep{
			id,
			StepForm,
			next,
		},
		fields,
	}
}

func NewChoiceStep(id string, next []Step) ChoiceStep {
	options := []string{}

	for _, s := range next {
		options = append(options, s.ID())
	}

	return ChoiceStep{
		BaseStep{
			id,
			StepChoice,
			next,
		},
		options,
	}
}
