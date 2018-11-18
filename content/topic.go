package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Topic struct {
	item.Item

	Name        string `json:"name"`
	Description string `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Topic within the CMS
// and implements editor.Editable
func (t *Topic) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(t,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Topic field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", t, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", t, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Topic editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Topic"] = func() interface{} { return new(Topic) }
}

// String defines how a Topic is printed. Update it using more descriptive
// fields from the Topic struct type
func (t *Topic) String() string {
	return fmt.Sprintf("Topic: %s", t.Name)
}
