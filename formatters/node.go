package formatters

type Status int

const (
	StatusUnchanged Status = iota
	StatusRemoved
	StatusAdded
	StatusChanged
	StatusNested
)

// Node is an internal representation of one difference item.
type Node struct {
	Key      string
	Status   Status
	OldValue any
	NewValue any
	Children []Node
}
