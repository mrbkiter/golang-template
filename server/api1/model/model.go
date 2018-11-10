package model

import (
	"encoding/json"

	"template.github.com/server/app"
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
func InternalErrorToAPIError(e *app.InternalError) *APIError {
	return &APIError{ErrorCode: convertErrorCode(&e.ErrorCode), ErrorMessage: &e.ErrorMessage}
}

func convertErrorCode(code *app.ErrorCode) ErrorCode {
	return ErrorCode(*code)
}

//APIErrorToJSON APIError to JSON
func APIErrorToJSON(err *APIError) string {
	result, _ := json.Marshal(err)
	return string(result)
}
