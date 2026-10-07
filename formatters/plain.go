package formatters

import (
	"fmt"
	"strings"
)

const complexValue = "[complex value]"

// Plain formats a diff tree as a list of changed properties.
func Plain(nodes []Node) string {
	return formatPlainNodes(nodes, "")
}

func formatPlainNodes(nodes []Node, parentPath string) string {
	lines := make([]string, 0, len(nodes))

	for _, node := range nodes {
		path := node.Key
		if parentPath != "" {
			path = parentPath + "." + node.Key
		}

		switch node.Status {
		case StatusNested:
			if line := formatPlainNodes(node.Children, path); line != "" {
				lines = append(lines, line)
			}
		case StatusRemoved:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", path))
		case StatusAdded:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", path, formatPlainValue(node.NewValue)))
		case StatusChanged:
			lines = append(lines, fmt.Sprintf(
				"Property '%s' was updated. From %s to %s",
				path,
				formatPlainValue(node.OldValue),
				formatPlainValue(node.NewValue),
			))
		}
	}

	return strings.Join(lines, "\n")
}

func formatPlainValue(value any) string {
	if _, ok := value.(map[string]any); ok {
		return complexValue
	}
	if stringValue, ok := value.(string); ok {
		return "'" + stringValue + "'"
	}
	if value == nil {
		return "null"
	}
	return fmt.Sprintf("%v", value)
}
