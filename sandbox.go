package dkr

import (
	"fmt"
	"os"
	"os/exec"
)

func Sandbox(image string, args... string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	env := []string{}
	for _, e := range os.Environ() {
		env = append(env, "-e" + e)
	}

	cmdArgs := []string{"run", "--rm", "-it", "-v=/var/run/docker.sock:/var/run/docker.sock",
		fmt.Sprintf("-v=%s:%s", "/", "/host"),
		fmt.Sprintf("-w=/host%s", pwd,)}
	cmdArgs = append(cmdArgs, env...)
	dockerConfig := append(cmdArgs, image)
	c := exec.Command("docker", append(dockerConfig, args...)...)
	c.Env = os.Environ()
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Run()
	return err
}
