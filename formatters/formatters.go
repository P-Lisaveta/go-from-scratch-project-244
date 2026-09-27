package formatters

import "fmt"

const (
	StylishFormat = "stylish"
	PlainFormat   = "plain"
)

// Format selects and applies a formatter to the diff tree.
func Format(nodes []Node, format string) (string, error) {
	switch format {
	case StylishFormat:
		return Stylish(nodes), nil
	case PlainFormat:
		return Plain(nodes), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}
