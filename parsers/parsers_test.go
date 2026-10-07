package parsers_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"code/parsers"
)

func TestParseJSON(t *testing.T) {
	path := filepath.Join("..", "testdata", "fixture", "file1.json")

	got, err := parsers.Parse(path)
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

func TestParseYAML(t *testing.T) {
	path := filepath.Join("..", "testdata", "fixture", "file1.yml")

	got, err := parsers.Parse(path)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if got["host"] != "hexlet.io" {
		t.Errorf("Parse() host = %v, want hexlet.io", got["host"])
	}
	if got["timeout"] != 50 {
		t.Errorf("Parse() timeout = %v, want 50", got["timeout"])
	}
	if got["follow"] != false {
		t.Errorf("Parse() follow = %v, want false", got["follow"])
	}
}

func TestParseYAMLExtension(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.yaml")
	content := "enabled: true\n"

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	got, err := parsers.Parse(path)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if got["enabled"] != true {
		t.Errorf("Parse() enabled = %v, want true", got["enabled"])
	}
}

func TestParseCaseInsensitiveExtension(t *testing.T) {
	jsonPath := filepath.Join(t.TempDir(), "config.JSON")
	yamlPath := filepath.Join(t.TempDir(), "config.YML")

	if err := os.WriteFile(jsonPath, []byte(`{"enabled": true}`), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if err := os.WriteFile(yamlPath, []byte("enabled: true\n"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	for _, path := range []string{jsonPath, yamlPath} {
		got, err := parsers.Parse(path)
		if err != nil {
			t.Fatalf("Parse(%s) error = %v", path, err)
		}
		if got["enabled"] != true {
			t.Errorf("Parse(%s) enabled = %v, want true", path, got["enabled"])
		}
	}
}

func TestParseInvalidFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	tests := []struct {
		name    string
		path    string
		content string
	}{
		{name: "invalid JSON", path: filepath.Join(dir, "invalid.json"), content: "{invalid}"},
		{name: "invalid YAML", path: filepath.Join(dir, "invalid.yaml"), content: "[invalid"},
		{name: "unsupported format", path: filepath.Join(dir, "invalid.txt"), content: "data"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := os.WriteFile(tt.path, []byte(tt.content), 0o600); err != nil {
				t.Fatalf("WriteFile() error = %v", err)
			}

			if _, err := parsers.Parse(tt.path); err == nil {
				t.Fatalf("Parse(%s) expected an error", tt.path)
			} else if !strings.Contains(err.Error(), tt.path) {
				t.Errorf("Parse(%s) error = %q, want it to contain the file path", tt.path, err)
			}
		})
	}
}
