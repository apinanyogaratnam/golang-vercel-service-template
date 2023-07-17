package models

type HealthCheckResponse struct {
	Status string `json:"status"`
	Load   Load   `json:"load"`
}

type Load struct {
	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
}
