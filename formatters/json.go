package formatters

import (
	"encoding/json"
)

type jsonNode struct {
	Key      string          `json:"key"`
	Type     string          `json:"type"`
	Value    json.RawMessage `json:"value,omitempty"`
	OldValue json.RawMessage `json:"oldValue,omitempty"`
	NewValue json.RawMessage `json:"newValue,omitempty"`
	Children []jsonNode      `json:"children,omitempty"`
}

// JSON formats a diff tree as a JSON array.
func JSON(nodes []Node) string {
	result, err := json.MarshalIndent(formatJSONNodes(nodes), "", "    ")
	if err != nil {
		// Node contains only JSON-compatible data, so marshalling must not fail.
		return "[]"
	}
	return string(result)
}

func formatJSONNodes(nodes []Node) []jsonNode {
	result := make([]jsonNode, 0, len(nodes))
	for _, node := range nodes {
		jsonItem := jsonNode{
			Key:  node.Key,
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

		result = append(result, jsonItem)
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
