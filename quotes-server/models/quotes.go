package models

// Quotes - the main struct for the quotes
// it contains the quote and the author
type Quotes struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}
