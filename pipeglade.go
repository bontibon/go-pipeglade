package pipeglade

import (
	"io"
	"os/exec"
	"strings"
	"sync"
)

type UI struct {
	C chan Event

	cmd *exec.Cmd

	stdout         io.ReadCloser
	stdoutReplacer *strings.Replacer

	outLock       sync.Mutex
	stdin         io.WriteCloser
	stdinReplacer *strings.Replacer
}

func Exec(interfaceFile string) (*UI, error) {
	proc := exec.Command("pipeglade", "-u", interfaceFile)
	stdin, err := proc.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := proc.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := proc.Start(); err != nil {
		return nil, err
	}
	ui := &UI{
		C: make(chan Event),

		cmd: proc,

		stdout:         stdout,
		stdoutReplacer: strings.NewReplacer(`\\`, `\`, `\r`, "\r", `\n`, "\n"),

		stdin:         stdin,
		stdinReplacer: strings.NewReplacer("\n", `\n`, "\r", `\r`, `\`, `\\`),
	}
	go ui.readRoutine()
	return ui, nil
}
