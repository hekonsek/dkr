package dkr

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type SandboxOptions struct {
	Out io.Writer
}

func Sandbox(image string, entrypoint []string, args []string, options *SandboxOptions) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	env := []string{}
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "HOME=") {
			homeSegments := strings.SplitN(e, "=", 2)
			e = fmt.Sprintf("HOME=/host%s", homeSegments[1])
		}
		env = append(env, "-e"+e)
	}

	cmdArgs := []string{"run", "--rm", "-v=/var/run/docker.sock:/var/run/docker.sock",
		fmt.Sprintf("-v=%s:%s", "/", "/host"),
		fmt.Sprintf("-w=/host%s", pwd)}
	isTty, err := isTty()
	if err != nil {
		return err
	}
	if isTty {
		cmdArgs = append(cmdArgs, "-it")
	}
	cmdArgs = append(cmdArgs, env...)
	dockerConfig := append(cmdArgs, image)
	if entrypoint != nil && len(entrypoint) > 0 {
		entrypointArg := fmt.Sprintf(`["%s"]`, strings.Join(entrypoint, `","`))
		dockerConfig = append(dockerConfig, "--entrypoint", entrypointArg)
	}
	c := exec.Command("docker", append(dockerConfig, args...)...)
	c.Env = os.Environ()
	c.Stdin = os.Stdin
	if options != nil && options.Out != nil {
		c.Stdout = options.Out
	} else {
		c.Stdout = os.Stdout
	}
	c.Stderr = os.Stderr
	err = c.Run()
	return err
}

func isTty() (bool, error) {
	cmd := exec.Command("docker", "run", "-it", "echo", "foo")
	out, err := cmd.CombinedOutput()
	if strings.Contains(string(out), "not a TTY") {
		return false, nil
	}
	return true, err
}
