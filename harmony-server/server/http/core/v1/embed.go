package v1

import "encoding/json"

type ActionType int
type ActionPresentation int
type FieldPresentation int

const (
	Normal ActionType = iota
	Primary
	Destructive

	Button ActionPresentation = iota
	Dropdown
	Menu
	SmallEntry
	LargeEntry

	Data FieldPresentation = iota
	CaptionedImage
	Row
)

type Action struct {
	Text string `json:"text,omitempty"`
	URL  string `json:"url,omitempty"`
	ID   string `json:"id,omitempty"`

	Type ActionType `json:"type,omitempty"`

	Children []*Action `json:"children,omitempty"`
}

type EmbedHeading struct {
	Text    string `json:"text,omitempty"`
	Subtext string `json:"subtext,omitempty"`
	URL     string `json:"url,omitempty"`
	Icon    string `json:"icon,omitempty"`
}

type EmbedField struct {
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Body     string `json:"body,omitempty"`

	ImageURL string `json:"image_url,omitempty"`

	Actions      []Action          `json:"actions,omitempty"`
	Presentation FieldPresentation `json:"presentation,omitempty"`
}

type Embed struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`

	Color int `json:"color,omitempty"`

	Header *EmbedHeading `json:"header,omitempty"`
	Footer *EmbedHeading `json:"footer,omitempty"`

	Fields []EmbedField `json:"fields,omitempty"`

	Actions []Action `json:"actions,omitempty"`
}

// CleanEmbed parses arbitrary JSON and returns known-good JSON
func CleanEmbed(in []byte) (out []byte, err error) {
	var e Embed
	err = json.Unmarshal(in, &e)
	if err != nil {
		return
	}
	out, err = json.Marshal(e)
	return
}

// CleanAction parses arbitrary JSON and returns known-good JSON
func CleanAction(in []byte) (out []byte, err error) {
	var a Action
	err = json.Unmarshal(in, &a)
	if err != nil {
		return
	}
	out, err = json.Marshal(a)
	return
}
