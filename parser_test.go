package code_test

import (
	"os"
	"path/filepath"
	"testing"

	"code"
)

func TestParseJSON(t *testing.T) {
	path := filepath.Join("testdata", "fixture", "file1.json")

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

func TestParseJSONCaseInsensitiveExtension(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.JSON")
	content := `{"enabled": true}`

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	got, err := code.Parse(path)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if got["enabled"] != true {
		t.Errorf("Parse() enabled = %v, want true", got["enabled"])
	}
}

func TestParseInvalidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")

	if err := os.WriteFile(path, []byte("{invalid}"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if _, err := code.Parse(path); err == nil {
		t.Fatal("Parse() expected an error for invalid JSON")
	}
}

func TestParseUnsupportedFormat(t *testing.T) {
	path := filepath.Join(t.TempDir(), "file.txt")

	if _, err := code.Parse(path); err == nil {
		t.Fatal("Parse() expected an error for unsupported format")
	}
}
