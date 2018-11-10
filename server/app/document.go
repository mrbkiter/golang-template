package app

import (
	"template.github.com/server/model"
)

type documentApp struct{}

func (docApp *documentApp) ParseResumeToCandidate(base64Content string) *model.Candidate {
	//url := config.Config().ParseResumeURL
	//apiKey := config.Config().ParseAPIKey
	return &model.Candidate{}
}
