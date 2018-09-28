package charlie

// Source structure contains data, that is used during parse process
// All of fields are optional and may not contain data
type Source struct {
	Title string // May contain title with release version and (but not necessary) release date
	Date  string // May contain release date
	Body  string // May contain issues collection to parse
}
