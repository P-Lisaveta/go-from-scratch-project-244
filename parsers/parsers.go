package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse reads a configuration file and returns its data as a map.
// The file format is determined by the file extension.
func Parse(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case ".json":
		return parseJSON(data)
	case ".yaml", ".yml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", filepath.Ext(path))
	}
}

func parseJSON(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON file: %w", err)
	}
	return result, nil
}

func parseYAML(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid YAML file: %w", err)
	}
	return result, nil
}
