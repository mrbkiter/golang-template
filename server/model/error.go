package model

import (
	"encoding/json"
	"fmt"
)

type ErrorCode string

const (
	ErrorCandidateNotFound  ErrorCode = "candidate_not_found"
	ErrorCandidateError     ErrorCode = "candidate_error"
	ErrorMissingCandidateId ErrorCode = "missing_candidate_id"
)

type InternalError struct {
	ErrorMessage string
	ErrorCode    ErrorCode
}

//ErrorFactory builder for app error
type ErrorFactory struct{}

func (e *InternalError) Error() string {
	return fmt.Sprintf("internal_business_error %v", e)
}

//BuildInternalError construct an internal error
func (ie *ErrorFactory) BuildInternalError(errorCode ErrorCode, msg string) *InternalError {
	var err *InternalError
	switch errorCode {
	case ErrorCandidateNotFound:
		err = &InternalError{ErrorCode: ErrorCandidateNotFound, ErrorMessage: msg}
		break
	case ErrorCandidateError:
		err = &InternalError{ErrorCode: ErrorCandidateError, ErrorMessage: msg}
		break
	case ErrorMissingCandidateId:
		err = &InternalError{ErrorCode: ErrorMissingCandidateId, ErrorMessage: msg}
		break
	}
	return err
}

//InternalErrorToJSON convert this object to Json string
func InternalErrorToJSON(e *InternalError) string {
	result, _ := json.Marshal(e)
	return string(result)
}
