package src

import "flag"

const (
	PROCFILE_ARG       = "p"
	PROCFILE_DEFAULT   = "Procfile"
	BACKGROUND_ARG     = "b"
	BACKGROUND_DEFAULT = false
)

type Args struct {
	Procfile   *string
	Background *bool
}

func ParseArguments() *Args {
	return &Args{
		Procfile:   flag.String(PROCFILE_ARG, PROCFILE_DEFAULT, "The Procfile to load"),
		Background: flag.Bool(BACKGROUND_ARG, BACKGROUND_DEFAULT, "Run the processes as a deamon"),
	}
}
