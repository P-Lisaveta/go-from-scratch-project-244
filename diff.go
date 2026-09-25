package code

import (
	"fmt"
	"reflect"
	"sort"

	"code/parsers"
)

const FormatStylish = "stylish"

type diffStatus int

const (
	statusUnchanged diffStatus = iota
	statusRemoved
	statusAdded
	statusChanged
	statusNested
)

type diffNode struct {
	key      string
	status   diffStatus
	oldValue any
	newValue any
	children []diffNode
}

// GenDiff returns a textual representation of the difference between
// two configuration files in the requested format. If format is empty,
// the stylish formatter is used.
func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := parsers.Parse(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := parsers.Parse(filepath2)
	if err != nil {
		return "", err
	}

	if format == "" {
		format = FormatStylish
	}

	diff := buildDiff(data1, data2)

	switch format {
	case FormatStylish:
		return formatStylish(diff), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}

func buildDiff(data1, data2 map[string]any) []diffNode {
	keys := make([]string, 0, len(data1)+len(data2))
	seen := make(map[string]struct{}, len(data1)+len(data2))
	for key := range data1 {
		if _, exists := seen[key]; !exists {
			keys = append(keys, key)
			seen[key] = struct{}{}
		}
	}
	for key := range data2 {
		if _, exists := seen[key]; !exists {
			keys = append(keys, key)
			seen[key] = struct{}{}
		}
	}
	sort.Strings(keys)

	nodes := make([]diffNode, 0, len(keys))
	for _, key := range keys {
		oldValue, inFirst := data1[key]
		newValue, inSecond := data2[key]
		oldMap, oldIsMap := oldValue.(map[string]any)
		newMap, newIsMap := newValue.(map[string]any)

		node := diffNode{key: key, oldValue: oldValue, newValue: newValue}

		switch {
		case inFirst && inSecond && oldIsMap && newIsMap:
			node.status = statusNested
			node.children = buildDiff(oldMap, newMap)
		case !inFirst:
			node.status = statusAdded
		case !inSecond:
			node.status = statusRemoved
		case reflect.DeepEqual(oldValue, newValue):
			node.status = statusUnchanged
		default:
			node.status = statusChanged
		}

		nodes = append(nodes, node)
	}

	return nodes
}
