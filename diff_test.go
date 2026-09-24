package code_test

import (
	"path/filepath"
	"testing"

	"code"
)

func TestGenDiffFlatJSON(t *testing.T) {
	filepath1 := filepath.Join("tests", "fixtures", "file1.json")
	filepath2 := filepath.Join("tests", "fixtures", "file2.json")

	got, err := code.GenDiff(filepath1, filepath2, "stylish")
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	want := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	if got != want {
		t.Errorf("GenDiff()\ngot:\n%s\nwant:\n%s", got, want)
	}
}
