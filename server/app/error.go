package app

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

type internalErrorBuilder struct{}

func (e *InternalError) Error() string {
	return fmt.Sprintf("internal_business_error %v", e)
}

func (ie *internalErrorBuilder) BuildInternalError(errorCode ErrorCode, msg string) *InternalError {
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
