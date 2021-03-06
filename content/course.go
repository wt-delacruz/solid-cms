package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Course struct {
	item.Item

	Title         string `json:"title"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	Author       string `json:"author"`
	Topic        string `json:"topic"`
}

// MarshalEditor writes a buffer of html to edit a Course within the CMS
// and implements editor.Editable
func (c *Course) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(c,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Course field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", c, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", c, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Requirements", c, map[string]string{
				"label":       "Requirements",
				"placeholder": "Enter the Requirements here",
			}),
		},
		editor.Field{
			View: editor.Input("Author", c, map[string]string{
				"label":       "Author",
				"type":        "text",
				"placeholder": "Enter the Author here",
			}),
		},
		editor.Field{
			View: reference.Select("Topic", c, map[string]string{
				"label": "Topic",
			},
				"Topic",
				`Topic: {{ .name }}`,
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Course editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Course"] = func() interface{} { return new(Course) }
}

// String defines how a Course is printed. Update it using more descriptive
// fields from the Course struct type
func (c *Course) String() string {
	return fmt.Sprintf("Course: %s", c.Title)
}
