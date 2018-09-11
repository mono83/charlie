package charlie

// Issue describes an issue, resolved or implemented in release
type Issue struct {
	ID         string
	Type       Type
	Components []string
	Message    string
}
