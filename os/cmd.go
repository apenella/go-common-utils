package common

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

// RunCmd takes a command and args and runs it, streaming output to stdout
func RunCmd(cmdName string, cmdArgs []string, prefix string) error {
	fmt.Printf(" Run: %s %s\n", cmdName, strings.Join(cmdArgs, " "))

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s =>  %s\n", prefix, scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
