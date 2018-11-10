package app

//app class that holds all service objects
type app struct {
	Config    *config
	Candidate *candidateApp
}

//App services object
var App = &app{Candidate: &candidateApp{},
	Config: &config{}}

//ErrorBuilder builder for error
var errorBuilder = &internalErrorBuilder{}
