package web

type WebResponse struct {
	Code   int         `json:"Code"`
	Status string      `json:"Status"`
	Data   interface{} `jsoon:"data,omitempty"`
}
