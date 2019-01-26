package model

// Issue describes an issue, resolved or implemented in release
type Issue struct {
	ID         int64
	IssueID    string
	Type       Type
	Components []string
	Message    string
}
