package formatters

import (
	"fmt"
	"sort"
	"strings"
)

const indentWidth = 4

// Stylish formats a diff tree in stylish representation.
func Stylish(nodes []Node) string {
	return "{\n" + formatStylishNodes(nodes, 1) + "\n}"
}

func formatStylishNodes(nodes []Node, depth int) string {
	lines := make([]string, 0, len(nodes))
	for _, node := range nodes {
		lines = append(lines, formatStylishNode(node, depth)...)
	}
	return strings.Join(lines, "\n")
}

func formatStylishNode(node Node, depth int) []string {
	switch node.Status {
	case StatusNested:
		return formatStylishNestedValue(" ", node.Key, node.Children, depth)
	case StatusRemoved:
		return formatStylishValue("-", node.Key, node.OldValue, depth)
	case StatusAdded:
		return formatStylishValue("+", node.Key, node.NewValue, depth)
	case StatusChanged:
		lines := formatStylishValue("-", node.Key, node.OldValue, depth)
		return append(lines, formatStylishValue("+", node.Key, node.NewValue, depth)...)
	default:
		return []string{formatStylishScalarLine(" ", node.Key, node.NewValue, depth)}
	}
}

func formatStylishValue(marker, key string, value any, depth int) []string {
	if childMap, ok := value.(map[string]any); ok {
		return formatStylishObject(marker, key, formatMapStylish(childMap, depth+1), depth)
	}
	return []string{formatStylishScalarLine(marker, key, value, depth)}
}

func formatStylishNestedValue(marker, key string, children []Node, depth int) []string {
	return formatStylishObject(marker, key, formatStylishNodes(children, depth+1), depth)
}

func formatStylishObject(marker, key, children string, depth int) []string {
	return []string{
		stylishPrefix(marker, depth) + key + ": {",
		children,
		strings.Repeat(" ", depth*indentWidth) + "}",
	}
}

func formatMapStylish(data map[string]any, depth int) string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		lines = append(lines, formatStylishValue(" ", key, data[key], depth)...)
	}
	return strings.Join(lines, "\n")
}

func formatStylishScalarLine(marker, key string, value any, depth int) string {
	return fmt.Sprintf("%s%s: %s", stylishPrefix(marker, depth), key, formatScalar(value))
}

func stylishPrefix(marker string, depth int) string {
	return strings.Repeat(" ", depth*indentWidth-2) + marker + " "
}

func formatScalar(value any) string {
	switch typedValue := value.(type) {
	case nil:
		return "null"
	case string:
		return typedValue
	default:
		return fmt.Sprintf("%v", typedValue)
	}
}
