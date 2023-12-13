package main

type FailureResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}
