package src

type ProcessType string

type ProcfileProcess struct {
	Type ProcessType
	Cmd  string
}

type Procfile struct {
	Processes []ProcfileProcess
}