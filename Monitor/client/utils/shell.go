package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecCommand(command string) (out string, err error) {
	commands := strings.Split(command, " ")
	cmd := exec.Command(commands[0], commands[1:]...)
	ret, err := cmd.Output()
	fmt.Printf("%s output is : %s", command, string(ret))
	return string(ret), err
}
