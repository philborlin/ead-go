package cmds

// KafkaCmds represents a Kafka command creator
type KafkaCmds struct {
	stack *Stack
}

// AppendCmd represents an append command
type AppendCmd struct {
	protobuf []byte
}

// Append puts a protobuf on the Kafka Cluster
func (c *KafkaCmds) Append(protobuf []byte) {
	c.stack.AddCmd <- AppendCmd{protobuf}
}

// NewKafkaCmds creates a KafkaCmds
func NewKafkaCmds(stack *Stack) *KafkaCmds {
	return &KafkaCmds{stack}
}
