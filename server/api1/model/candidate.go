package model

import (
	"encoding/json"
	"io"

	imodel "template.github.com/server/model"
)

//Candidate candidate type
type Candidate struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//CandidateToInternalCandidate convert candidate v1 -> internal candidate
func CandidateToInternalCandidate(candidate *Candidate) *imodel.Candidate {
	ic := &imodel.Candidate{ID: candidate.ID, FirstName: candidate.FirstName, LastName: candidate.LastName}
	return ic
}

//InternalCandidateToCandidate convert Internal Candidate to Candidate v1
func InternalCandidateToCandidate(iCandidate *imodel.Candidate) *Candidate {
	candidate := &Candidate{ID: iCandidate.ID, FirstName: iCandidate.FirstName, LastName: iCandidate.LastName}
	return candidate
}

//CandidateToJSON marshal candidate to JSON String
func CandidateToJSON(candidate *Candidate) string {
	b, _ := json.Marshal(candidate)
	return string(b)
}

//CandidateFromJSON read json string in byte[], then decode to Candidate Object
func CandidateFromJSON(data io.Reader) *Candidate {
	var candidate *Candidate
	json.NewDecoder(data).Decode(&candidate)
	return candidate
}

//CandidateListToJSON candidate list to json string
func CandidateListToJSON(r []*Candidate) string {
	b, _ := json.Marshal(r)
	return string(b)
}

//CandidateListFromJSON from JSON string to list of candidates
func CandidateListFromJSON(data io.Reader) []*Candidate {
	var candidates []*Candidate
	json.NewDecoder(data).Decode(&candidates)
	return candidates
}
