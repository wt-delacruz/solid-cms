package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Activity struct {
	item.Item

	Title        string `json:"title"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	Type         string `json:"type"`
	Data         string `json:"data"`
}

// MarshalEditor writes a buffer of html to edit a Activity within the CMS
// and implements editor.Editable
func (a *Activity) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Activity field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", a, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", a, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Instructions", a, map[string]string{
				"label":       "Instructions",
				"placeholder": "Enter the Instructions here",
			}),
		},
		editor.Field{
			View: editor.Input("Type", a, map[string]string{
				"label":       "Type",
				"type":        "text",
				"placeholder": "Enter the Type here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Data", a, map[string]string{
				"label":       "Data",
				"placeholder": "Enter the Data here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Activity editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Activity"] = func() interface{} { return new(Activity) }
}

// String defines how a Activity is printed. Update it using more descriptive
// fields from the Activity struct type
func (a *Activity) String() string {
	return fmt.Sprintf("Activity: %s", a.UUID)
}
