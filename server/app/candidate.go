package app

import (
	ctx1 "context"
	"log"

	"template.github.com/server/repo"

	"template.github.com/server/model"
)

type context = ctx1.Context

//CandidateApp candidate service
type candidateApp struct {
}

//CreateCandidate function to create candidate
func (c *candidateApp) CreateCandidate(ctx *context, candidate *model.Candidate) (string, *model.InternalError) {
	log.Printf("CandidateApp %v\n", candidate)
	if candidate.ID != "" {
		log.Printf("You can not specify candidate id. It should be auto generated from system %v.", candidate)
		return "", errorFactory.BuildInternalError(model.ErrorCandidateError, "")
	}
	id, error := repo.Repo().CreateCandidate(ctx, candidate)
	log.Printf("candidate %v written successfully\n", candidate)
	return id, error
}

//FindCandidateByID find a candidate by id. Input is candidateId (not empty)
func (c *candidateApp) FindCandidateByID(ctx *context, candidateID string) (*model.Candidate, *model.InternalError) {
	if candidateID == "" {
		return nil, errorFactory.BuildInternalError(model.ErrorCandidateNotFound, "")
	}
	candidate := repo.Repo().FindCandidateByID(ctx, candidateID)
	if candidate == nil {
		return nil, errorFactory.BuildInternalError(model.ErrorCandidateNotFound, "")
	}
	return candidate, nil
}
