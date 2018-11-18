package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Skill struct {
	item.Item

	Name        string `json:"name"`
	Description string `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Skill within the CMS
// and implements editor.Editable
func (s *Skill) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(s,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Skill field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", s, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", s, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Skill editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Skill"] = func() interface{} { return new(Skill) }
}

// String defines how a Skill is printed. Update it using more descriptive
// fields from the Skill struct type
func (s *Skill) String() string {
	return fmt.Sprintf("Skill: %s", s.Name)
}
