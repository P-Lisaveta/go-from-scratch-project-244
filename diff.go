package code

import (
	"fmt"
	"sort"
	"strings"

	"code/parsers"
)

// GenDiff returns a textual representation of the difference between
// two configuration files in the requested format.
func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := parsers.Parse(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := parsers.Parse(filepath2)
	if err != nil {
		return "", err
	}

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

	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		oldValue, inFirst := data1[key]
		newValue, inSecond := data2[key]

		switch {
		case inFirst && inSecond:
			if oldValue == newValue {
				lines = append(lines, diffLine(" ", key, oldValue))
				continue
			}
			lines = append(lines, diffLine("-", key, oldValue))
			lines = append(lines, diffLine("+", key, newValue))
		case inFirst:
			lines = append(lines, diffLine("-", key, oldValue))
		case inSecond:
			lines = append(lines, diffLine("+", key, newValue))
		}
	}

	return "{\n" + strings.Join(lines, "\n") + "\n}", nil
}

func diffLine(marker, key string, value any) string {
	return fmt.Sprintf("  %s %s: %s", marker, key, formatValue(value))
}

func formatValue(value any) string {
	if stringValue, ok := value.(string); ok {
		return stringValue
	}
	return fmt.Sprintf("%v", value)
}
