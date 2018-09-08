package clipboard

import (
	"fmt"
	"os/exec"
)

func Pop() (content string, err error) {
	command := exec.Command("pbpaste")
	c, err := command.Output()
	if err != nil {
		return "", err
	}
	return string(c), nil
}

func Push(path string) error {

	cmdText := fmt.Sprintf("cat %s | pbcopy", path)
	err := exec.Command("sh", "-c", cmdText).Run()
	return err
}
