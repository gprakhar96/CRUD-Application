package errors

import "net/http"

type RestErrors struct {
	Message string `json:"message"`
	Status  int    `json:"status_code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "INTERNAL_SERVER_ERROR",
	}
}

func New404Error(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "404_RESOURCE_NOT_FOUND",
	}
}
