package app

import (
	"log"
	"strings"

	"github.com/rs/xid"
	"template.github.com/server/model"
)

//CandidateApp candidate service
type candidateApp struct {
}

//CreateCandidate function to create candidate
func (c *candidateApp) CreateCandidate(candidate *model.Candidate) (string, *InternalError) {
	log.Printf("CandidateApp %v\n", candidate)
	if candidate.ID == "1234" {
		log.Printf("error called candidate %v.", candidate)
		return "", errorBuilder.BuildInternalError(ErrorCandidateNotFound, "")
	}
	log.Printf("candidate %v written successfully\n", candidate)
	return xid.New().String(), nil
}

//FindCandidateByID find a candidate by id. Input is candidateId (not empty)
func (c *candidateApp) FindCandidateByID(candidateID string) (*model.Candidate, *InternalError) {
	if candidateID == "" || strings.Compare(candidateID, "1234") == 0 {
		return nil, errorBuilder.BuildInternalError(ErrorCandidateNotFound, "")
	}
	candidate := &model.Candidate{ID: candidateID, FirstName: "First Name", LastName: "Last Name"}
	return candidate, nil
}
