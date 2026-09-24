package code_test

import (
	"path/filepath"
	"testing"

	"code"
)

func TestParseJSON(t *testing.T) {
	path := filepath.Join("tests", "fixtures", "file1.json")

	got, err := code.Parse(path)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if got["host"] != "hexlet.io" {
		t.Errorf("Parse() host = %v, want hexlet.io", got["host"])
	}
	if got["timeout"] != float64(50) {
		t.Errorf("Parse() timeout = %v, want 50", got["timeout"])
	}
	if got["follow"] != false {
		t.Errorf("Parse() follow = %v, want false", got["follow"])
	}
}

func TestParseUnsupportedFormat(t *testing.T) {
	path := filepath.Join(t.TempDir(), "file.txt")

	if _, err := code.Parse(path); err == nil {
		t.Fatal("Parse() expected an error for unsupported format")
	}
}
