package executor

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type CommandExecutor interface {
	Execute(ctx context.Context, command string, workDir string) error
}

type shellCommandExecutor struct{}

func NewShellCommandExecutor() CommandExecutor {
	return &shellCommandExecutor{}
}

func (e *shellCommandExecutor) Execute(ctx context.Context, command string, workDir string) error {
	log.Println("Executing command:", command)

	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.CommandContext(ctx, "sh", "-c", command)
	} else {
		cmd = exec.CommandContext(ctx, "cmd", "/C", command)
	}

	var stderrBuf bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderrBuf
	if workDir != "" {
		cmd.Dir = workDir
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %w: %s", err, stderrBuf.String())
	}

	return nil
}
