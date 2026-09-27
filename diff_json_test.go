package code_test

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"code"
)

func TestGenDiffJSON(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.json")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.json")

	got, err := code.GenDiff(filepath1, filepath2, code.FormatJSON)
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	if !json.Valid([]byte(got)) {
		t.Fatalf("GenDiff() returned invalid JSON:\n%s", got)
	}

	var nodes map[string]map[string]any
	if err := json.Unmarshal([]byte(got), &nodes); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if len(nodes) != 4 {
		t.Fatalf("GenDiff() returned %d root nodes, want 4", len(nodes))
	}

	for _, key := range []string{"common", "group1", "group2", "group3"} {
		if _, exists := nodes[key]; !exists {
			t.Fatalf("GenDiff() does not contain root node %q", key)
		}
	}

	common := nodes["common"]
	if common["type"] != "nested" {
		t.Fatalf("unexpected common node: %#v", common)
	}

	children, ok := common["children"].(map[string]any)
	if !ok {
		t.Fatalf("common.children is not an object: %#v", common["children"])
	}
	if len(children) != 7 {
		t.Fatalf("common.children has %d items, want 7", len(children))
	}

	group2 := nodes["group2"]
	if group2["type"] != "removed" {
		t.Fatalf("unexpected group2 node: %#v", group2)
	}
	if _, ok := group2["value"].(map[string]any); !ok {
		t.Fatalf("group2.value is not an object: %#v", group2["value"])
	}

	group3 := nodes["group3"]
	if group3["type"] != "added" {
		t.Fatalf("unexpected group3 node: %#v", group3)
	}
	if _, ok := group3["value"].(map[string]any); !ok {
		t.Fatalf("group3.value is not an object: %#v", group3["value"])
	}
}

func TestGenDiffJSONYAML(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.yml")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.yml")

	jsonFromJSON, err := code.GenDiff(filepath.Join("testdata", "fixture", "nested1.json"), filepath.Join("testdata", "fixture", "nested2.json"), code.FormatJSON)
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	got, err := code.GenDiff(filepath1, filepath2, code.FormatJSON)
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	if got != jsonFromJSON {
		t.Errorf("GenDiff() for YAML\ngot:\n%s\nwant:\n%s", got, jsonFromJSON)
	}
}
