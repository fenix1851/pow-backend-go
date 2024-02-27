package requests

type PoWData struct {
	Data              int64  `json:"data"`
	LeadingZerosCount int64  `json:"leadingZerosCount"`
	Token             string `json:"token"`
}
