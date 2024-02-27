package models

// Tokens - the main struct for the pow
// it contains the token, the data and the timestamp
type Tokens struct {
	Token             string `json:"token"`
	Data              int64  `json:"data"`
	LeadingZerosCount int64  `json:"leadingZerosCount"`
	Timestamp         int64  `json:"timestamp"`
}
