package model

import (
	"encoding/json"
	"net/http"

	"template.github.com/server/model"
)

type ErrorCode string

const (
	CandidateNotFound ErrorCode = "candidate_not_found"
	CandidateError    ErrorCode = ""
)

type APIError struct {
	ErrorMessage *string   `json:"message"`
	ErrorCode    ErrorCode `json:"code"`
}

//InternalErrorToAPIError internal error to api error
func InternalErrorToAPIError(e *model.InternalError) (*APIError, int) {
	return &APIError{ErrorCode: convertErrorCode(&e.ErrorCode), ErrorMessage: &e.ErrorMessage}, http.StatusBadRequest
}

func convertErrorCode(code *model.ErrorCode) ErrorCode {
	return ErrorCode(*code)
}

//APIErrorToJSON APIError to JSON
func APIErrorToJSON(err *APIError) string {
	result, _ := json.Marshal(err)
	return string(result)
}
