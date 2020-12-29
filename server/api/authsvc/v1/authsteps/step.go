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
	CanGoBack() bool
	SetPreviousStep(s Step)
	GetPreviousStep() Step
	StepType() StepType
	SubSteps() []Step
}

// BaseStep is the base implementation for the Step interface
type BaseStep struct {
	id           string
	canGoBack    bool
	stepType     StepType
	subSteps     []Step
	previousStep Step
}

func (b *BaseStep) ID() string {
	return b.id
}

func (b *BaseStep) CanGoBack() bool {
	return b.canGoBack
}

func (b BaseStep) StepType() StepType {
	return b.stepType
}

func (b *BaseStep) SubSteps() []Step {
	return b.subSteps
}

func (b *BaseStep) AddStep(s Step) {
	b.subSteps = append(b.subSteps, s)
}

func (b *BaseStep) SetPreviousStep(s Step) {
	b.previousStep = s
}

func (b *BaseStep) GetPreviousStep() Step {
	return b.previousStep
}

// FormField ...
type FormField struct {
	Name      string
	FieldType string
}

// FormStep ...
type FormStep struct {
	*BaseStep
	Fields []FormField
}

// ChoiceStep ...
type ChoiceStep struct {
	*BaseStep
	Choices []string
}

func NewFormStep(id string, canGoBack bool, fields []FormField, next []Step) *FormStep {
	return &FormStep{
		&BaseStep{
			id,
			canGoBack,
			StepForm,
			next,
			nil,
		},
		fields,
	}
}

func NewChoiceStep(id string, canGoBack bool, next []Step) *ChoiceStep {
	options := []string{}

	for _, s := range next {
		options = append(options, s.ID())
	}

	return &ChoiceStep{
		&BaseStep{
			id,
			canGoBack,
			StepChoice,
			next,
			nil,
		},
		options,
	}
}
