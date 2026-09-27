package formatters

import (
	"encoding/json"
)

type jsonNode struct {
	Type     string              `json:"type"`
	Value    json.RawMessage     `json:"value,omitempty"`
	OldValue json.RawMessage     `json:"oldValue,omitempty"`
	NewValue json.RawMessage     `json:"newValue,omitempty"`
	Children map[string]jsonNode `json:"children,omitempty"`
}

// JSON formats a diff tree as a JSON object. Keys are property names.
func JSON(nodes []Node) string {
	result, err := json.MarshalIndent(formatJSONNodes(nodes), "", "    ")
	if err != nil {
		// Node contains only JSON-compatible data, so marshalling must not fail.
		return "{}"
	}
	return string(result)
}

func formatJSONNodes(nodes []Node) map[string]jsonNode {
	result := make(map[string]jsonNode, len(nodes))
	for _, node := range nodes {
		jsonItem := jsonNode{
			Type: formatJSONType(node.Status),
		}

		switch node.Status {
		case StatusNested:
			jsonItem.Children = formatJSONNodes(node.Children)
		case StatusAdded, StatusUnchanged:
			jsonItem.Value = marshalJSONValue(node.NewValue)
		case StatusRemoved:
			jsonItem.Value = marshalJSONValue(node.OldValue)
		case StatusChanged:
			jsonItem.OldValue = marshalJSONValue(node.OldValue)
			jsonItem.NewValue = marshalJSONValue(node.NewValue)
		}

		result[node.Key] = jsonItem
	}
	return result
}

func formatJSONType(status Status) string {
	switch status {
	case StatusNested:
		return "nested"
	case StatusRemoved:
		return "removed"
	case StatusAdded:
		return "added"
	case StatusChanged:
		return "changed"
	default:
		return "unchanged"
	}
}

func marshalJSONValue(value any) json.RawMessage {
	if value == nil {
		return json.RawMessage("null")
	}

	result, err := json.Marshal(value)
	if err != nil {
		return json.RawMessage("null")
	}
	return result
}
