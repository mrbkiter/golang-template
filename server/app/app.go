package app

import "template.github.com/server/model"

//app class that holds all service objects
type app struct {
	Candidate *candidateApp
}

//App services object
var App = &app{Candidate: &candidateApp{}}

//ErrorBuilder builder for error
var errorFactory = &model.ErrorFactory{}
