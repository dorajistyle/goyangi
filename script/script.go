package script

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func RunScript(commands []string) {
	entireScript := strings.NewReader(strings.Join(commands, "\n"))
	bash := exec.Command(Bash)
	stdin, _ := bash.StdinPipe()
	stdout, _ := bash.StdoutPipe()
	stderr, _ := bash.StderrPipe()

	wait := sync.WaitGroup{}
	wait.Add(3)
	go func() {
		io.Copy(stdin, entireScript)
		stdin.Close()
		wait.Done()
	}()
	go func() {
		io.Copy(os.Stdout, stdout)
		wait.Done()
	}()
	go func() {
		io.Copy(os.Stderr, stderr)
		wait.Done()
	}()

	bash.Start()
	wait.Wait()
	bash.Wait()
}
