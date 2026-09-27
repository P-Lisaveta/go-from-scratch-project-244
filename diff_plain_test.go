package code_test

import (
	"path/filepath"
	"testing"

	"code"
)

func TestGenDiffPlainJSON(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.json")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.json")

	got, err := code.GenDiff(filepath1, filepath2, code.FormatPlain)
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	want := `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`

	if got != want {
		t.Errorf("GenDiff()\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestGenDiffPlainYAML(t *testing.T) {
	filepath1 := filepath.Join("testdata", "fixture", "nested1.yml")
	filepath2 := filepath.Join("testdata", "fixture", "nested2.yml")

	got, err := code.GenDiff(filepath1, filepath2, code.FormatPlain)
	if err != nil {
		t.Fatalf("GenDiff() error = %v", err)
	}

	want := `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`

	if got != want {
		t.Errorf("GenDiff()\ngot:\n%s\nwant:\n%s", got, want)
	}
}
