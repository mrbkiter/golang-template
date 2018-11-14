package repo

import (
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

//FindCandidateByID find candidate by id. Return nil if candidate not found
func (r *Repository) FindCandidateByID(ctx *context, candidateID string) *model.Candidate {
	db := r.openConnection(ctx)
	candidate := &Candidate{}
	db.Where("id = ?", candidateID).First(candidate)
	if candidate.ID == "" {
		return nil
	}
	return candidate.toCandidateModel()
}

//CreateCandidate create candidate. return id if done. otherwise internal error
func (r *Repository) CreateCandidate(ctx *context, candidate *model.Candidate) (string, *model.InternalError) {
	//db := r.openConnection(ctx)
	//add your code here

	//check error if exists

	//return
	return "1234", nil
}
