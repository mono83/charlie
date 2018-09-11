package charlie

// Type describes issue type
type Type byte

// List of defined issue types
const (
	Notice      Type = 0
	Feature     Type = 1
	BugFix      Type = 11
	SecurityFix Type = 12
)
