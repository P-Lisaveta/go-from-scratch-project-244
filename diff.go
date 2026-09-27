package code

import (
	"reflect"
	"sort"

	"code/formatters"
	"code/parsers"
)

const (
	FormatStylish = formatters.StylishFormat
	FormatPlain   = formatters.PlainFormat
	FormatJSON    = formatters.JSONFormat
)

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
		format = formatters.StylishFormat
	}

	return formatters.Format(buildDiff(data1, data2), format)
}

func buildDiff(data1, data2 map[string]any) []formatters.Node {
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

	nodes := make([]formatters.Node, 0, len(keys))
	for _, key := range keys {
		oldValue, inFirst := data1[key]
		newValue, inSecond := data2[key]
		oldMap, oldIsMap := oldValue.(map[string]any)
		newMap, newIsMap := newValue.(map[string]any)

		node := formatters.Node{Key: key, OldValue: oldValue, NewValue: newValue}

		switch {
		case inFirst && inSecond && oldIsMap && newIsMap:
			node.Status = formatters.StatusNested
			node.Children = buildDiff(oldMap, newMap)
		case !inFirst:
			node.Status = formatters.StatusAdded
		case !inSecond:
			node.Status = formatters.StatusRemoved
		case reflect.DeepEqual(oldValue, newValue):
			node.Status = formatters.StatusUnchanged
		default:
			node.Status = formatters.StatusChanged
		}

		nodes = append(nodes, node)
	}

	return nodes
}
