package handler

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

var invalidParams = &ErrorResponse{ErrorMessage: "Invalid params."}
