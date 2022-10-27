package src

import (
	"context"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type Runner struct {
	ProcfileProcess ProcfileProcess
	Args			*Args
	ctx 			context.Context
	update 			chan RunnerUpdateResult
	exitChan 		chan RunnerExitResult
}

type RunnerExitResult struct {
	Error error
	ExitCode int
}

type RunnerUpdateResult struct {
	Error error
	Pid int
}

func NewRunner(
	p ProcfileProcess,
	a *Args, 
	ctx context.Context,
	update chan RunnerUpdateResult,
	exitChan chan RunnerExitResult,
	) *Runner {
	return &Runner{
		p, a, ctx,  update, exitChan,
	}
}


func (runner *Runner)Run() { 
	cmds := strings.Fields(runner.ProcfileProcess.Cmd)
	procFileDir, err := filepath.Abs(filepath.Dir(*runner.Args.Procfile))
    if err != nil {
		runner.update <- RunnerUpdateResult{err, -1}
		// cancel the context
		(runner.ctx).Done()
		return
    }

	subProcess := exec.CommandContext(runner.ctx,cmds[0], cmds[1:]...)

	// Set the working directory to the procfile directory
	subProcess.Dir = procFileDir

	if *runner.Args.Background {
		// Run the process in the background
		subProcess.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}

	// Start the process
	err = subProcess.Start()
	if err != nil {
		runner.update <- RunnerUpdateResult{err, -1}
		// cancel the context
		(runner.ctx).Done()
		return
	}

	// Send the pid to the update channel
	runner.update <- RunnerUpdateResult{nil, subProcess.Process.Pid}

	// Wait for the process to exit
	err = subProcess.Wait()
	if err != nil {
		runner.update <- RunnerUpdateResult{err, -1}
		return
	}

	// Send the exit code to the exit channel
	runner.exitChan <- RunnerExitResult{nil, subProcess.ProcessState.ExitCode()}

}
