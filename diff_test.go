package code_test

import (
	"path/filepath"
	"testing"

	"code"
)

func TestGenDiffNestedJSON(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.json")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.json")

	assertGenDiff(t, filepath1, filepath2, nestedStylishDiff)
}

func TestGenDiffNestedYAML(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.yml")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.yml")

	assertGenDiff(t, filepath1, filepath2, nestedStylishDiff)
}

func TestGenDiffDefaultFormat(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.json")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.json")

	got, err := code.GenDiff(filepath1, filepath2, "")
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	if got != nestedStylishDiff {
		t.Errorf("GenDiff()\ngot:\n%s\nwant:\n%s", got, nestedStylishDiff)
	}
}

func TestGenDiffUnsupportedFormat(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.json")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.json")

	if _, err := code.GenDiff(filepath1, filepath2, "unknown"); err == nil {
		t.Fatal("GenDiff() expected an error for unsupported format")
	}
}

func assertGenDiff(t *testing.T, filepath1, filepath2, want string) {
	t.Helper()

	got, err := code.GenDiff(filepath1, filepath2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	if got != want {
		t.Errorf("GenDiff()\ngot:\n%s\nwant:\n%s", got, want)
	}
}

const nestedStylishDiff = `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: 
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`
