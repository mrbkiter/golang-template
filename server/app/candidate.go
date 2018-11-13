package app

import (
	ctx1 "context"
	"log"

	"template.github.com/server/repo"

	"github.com/rs/xid"
	"template.github.com/server/model"
)

type context = ctx1.Context

//CandidateApp candidate service
type candidateApp struct {
}

//CreateCandidate function to create candidate
func (c *candidateApp) CreateCandidate(ctx *context, candidate *model.Candidate) (string, *InternalError) {
	log.Printf("CandidateApp %v\n", candidate)
	if candidate.ID == "1234" {
		log.Printf("error called candidate %v.", candidate)
		return "", errorBuilder.BuildInternalError(ErrorCandidateNotFound, "")
	}
	log.Printf("candidate %v written successfully\n", candidate)
	return xid.New().String(), nil
}

//FindCandidateByID find a candidate by id. Input is candidateId (not empty)
func (c *candidateApp) FindCandidateByID(ctx *context, candidateID string) (*model.Candidate, *InternalError) {
	if candidateID == "" {
		return nil, errorBuilder.BuildInternalError(ErrorCandidateNotFound, "")
	}
	candidate := repo.Repo().FindCandidateByID(ctx, candidateID)
	if candidate == nil || candidate.ID == "" {
		return nil, errorBuilder.BuildInternalError(ErrorCandidateNotFound, "")
	}
	return candidate, nil
}
