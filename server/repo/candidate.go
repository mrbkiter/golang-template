package repo

import (
	"context"

	"template.github.com/server/model"
)

type Candidate struct {
	ID        string `gorm:"column:id;primary_key`
	FirstName string
	LastName  string
}

func (Candidate) TableName() string {
	return "candidate"
}

func (candidate *Candidate) toCandidateModel() *model.Candidate {
	mCandidate := &model.Candidate{}
	mCandidate.FirstName = candidate.FirstName
	mCandidate.LastName = candidate.LastName
	mCandidate.ID = candidate.ID
	return mCandidate
}

//FindCandidateByID find candidate by id
func (r *Repository) FindCandidateByID(ctx *context.Context, candidateID string) *model.Candidate {
	db := r.openConnection(ctx)
	candidate := &Candidate{}
	db.Where("id = ?", candidateID).First(candidate)

	return candidate.toCandidateModel()
}
