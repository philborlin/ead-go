package cmds

// SQLCmds represents a Kafka command creator
type SQLCmds struct {
	stack *Stack
}

// InsertCmd represents an insert command
type InsertCmd struct {
	sql    string
	values []string
}

// Insert inserts into the db
func (c *SQLCmds) Insert(sql string, values ...string) int {
	i := new(int)
	c.stack.AddCmd <- InsertCmd{sql, values}
	return *i
}

// NewSQLCmds creates a SQLCmds
func NewSQLCmds(stack *Stack) *SQLCmds {
	return &SQLCmds{stack}
}
