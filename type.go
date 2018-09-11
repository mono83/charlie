package charlie

type Type byte

const (
	Notice      Type = 0
	Feature     Type = 1
	BugFix      Type = 11
	SecurityFix Type = 12
)
