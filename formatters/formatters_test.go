package formatters_test

import (
	"testing"

	"code/formatters"
)

func TestFormat(t *testing.T) {
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
	}

	stylish, err := formatters.Format(nodes, formatters.StylishFormat)
	if err != nil {
		t.Fatalf("Format() error = %v", err)
	}

	wantStylish := `{
    nested: {
      + added: value
    }
  - changed: old
  + changed: 42
}`

	if stylish != wantStylish {
		t.Errorf("Format() stylish\ngot:\n%s\nwant:\n%s", stylish, wantStylish)
	}

	plain, err := formatters.Format(nodes, formatters.PlainFormat)
	if err != nil {
		t.Fatalf("Format() error = %v", err)
	}

	wantPlain := "Property 'nested.added' was added with value: 'value'\nProperty 'changed' was updated. From 'old' to 42"

	if plain != wantPlain {
		t.Errorf("Format() plain\ngot:\n%s\nwant:\n%s", plain, wantPlain)
	}
}

func TestFormatUnsupported(t *testing.T) {
	if _, err := formatters.Format(nil, "unknown"); err == nil {
		t.Fatal("Format() expected an error for unsupported format")
	}
}
