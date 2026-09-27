package formatters_test

import (
	"testing"

	"code/formatters"
)

func TestJSON(t *testing.T) {
	nodes := []formatters.Node{
		{
			Key:    "nested",
			Status: formatters.StatusNested,
			Children: []formatters.Node{
				{Key: "added", Status: formatters.StatusAdded, NewValue: "value"},
			},
		},
		{
			Key:      "changed",
			Status:   formatters.StatusChanged,
			OldValue: "old",
			NewValue: 42,
		},
		{
			Key:      "removed",
			Status:   formatters.StatusRemoved,
			OldValue: nil,
		},
	}

	got, err := formatters.Format(nodes, formatters.JSONFormat)
	if err != nil {
		t.Fatalf("Format() error = %v", err)
	}

	want := `[
    {
        "key": "nested",
        "type": "nested",
        "children": [
            {
                "key": "added",
                "type": "added",
                "value": "value"
            }
        ]
    },
    {
        "key": "changed",
        "type": "changed",
        "oldValue": "old",
        "newValue": 42
    },
    {
        "key": "removed",
        "type": "removed",
        "value": null
    }
]`

	if got != want {
		t.Errorf("Format() JSON\ngot:\n%s\nwant:\n%s", got, want)
	}
}
