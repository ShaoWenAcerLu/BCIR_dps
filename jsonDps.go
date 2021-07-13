package dps

// Response
type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"Status"`
}

type SuccessResponse struct {
	Messages []string `json:"Message"`
	Status   string   `json:"Status"`
}
