package code_test

import (
	"os"
	"path/filepath"
	"testing"

	"code"
)

func TestGenDiffFlatJSON(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "file1.json")
	filepath2 := filepath.Join("testdata", "fixture", "file2.json")

	assertGenDiff(t, filepath1, filepath2, `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`)
}

func TestGenDiffFlatYAML(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "file1.yml")
	filepath2 := filepath.Join("testdata", "fixture", "file2.yml")

	assertGenDiff(t, filepath1, filepath2, `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`)
}

func TestGenDiffFlatJSONCases(t *testing.T) {
	dir := t.TempDir()
	filepath1 := writeTestFile(t, dir, "first.json", `{
		"changed": "before",
		"common": "same",
		"removed": 42
	}`)
	filepath2 := writeTestFile(t, dir, "second.json", `{
		"added": true,
		"changed": "after",
		"common": "same"
	}`)

	assertGenDiff(t, filepath1, filepath2, `{
  + added: true
  - changed: before
  + changed: after
    common: same
  - removed: 42
}`)
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

func writeTestFile(t *testing.T, dir, name, content string) string {
	t.Helper()

	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	return path
}
