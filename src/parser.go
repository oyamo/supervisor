package src

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Parser struct {
   Procfile string
}


func NewParser(p string) *Parser {
	return &Parser{p}
}

//CheckProcfile Check if the procfile exists and return it
func (p *Parser) GetProcfile() ([] string, error){
	file, err := os.Open(p.Procfile)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	// Read the file

	fileReader := bufio.NewReader(file)
	procfileContents := make([]string, 0)

	for {
		line, err := fileReader.ReadString ('\n')
		if err != nil {
			break
		}
		procfileContents = append(procfileContents, line)
	}

	return procfileContents, nil
	
}

func (p *Parser) Parse() ([]ProcfileProcess, error) {
	procfileProcesses := make([]ProcfileProcess, 0)
    procfileContents, err := p.GetProcfile()
	if err != nil {
		return nil, err
	}

	for lc, line := range procfileContents {
		lastIndexofColon := strings.LastIndex(line, ":")
		if lastIndexofColon == -1 {
			errorMessage := fmt.Sprintf("Line %d of the procfile does not contain a colon\n", lc)
			errorMessage += fmt.Sprintf("%d | %s", lc, line)
			return nil, errors.New(errorMessage)
		}

		// Check process type
		processType := strings.TrimSpace(line[:lastIndexofColon])
		if len(processType) == 0 {
			errorMessage := fmt.Sprintf("Line %d of the procfile does not contain a process type\n", lc)
			errorMessage += fmt.Sprintf("%d | %s%s%s", lc, WARN,line,END)
			return nil, errors.New(errorMessage)
		}

		if processType != "web" && processType != "command" {
			errorMessage := fmt.Sprintf("Line %d of the procfile contains an invalid process type\n", lc)
			errorMessage += fmt.Sprintf("%d | %s%s%s", lc, WARN,line,END)
			return nil, errors.New(errorMessage)
		}
			

		// Check command
		command := strings.TrimSpace(line[lastIndexofColon+1:])
		if len(command) == 0 {
			errorMessage := fmt.Sprintf("Line %d of the procfile does not contain a command\n", lc)
			errorMessage += fmt.Sprintf("%d | %s%s%s", lc, WARN, line, END)
			return nil, errors.New(errorMessage)
		}

		procfileProcesses = append(procfileProcesses, ProcfileProcess{
			Type: ProcessType(processType),
			Cmd: command,
		})
	}

	return procfileProcesses, nil
}