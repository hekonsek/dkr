package dkr

import (
	"io/ioutil"
)

func CopyProxy(command string) (target string, err error) {
	proxyBytes, err := ioutil.ReadFile("/usr/bin/dkr-proxy")
	if err != nil {
		return "", err
	}
	target = "/usr/bin/" + command
	err = ioutil.WriteFile(target, proxyBytes, 0555)
	if err != nil {
		return "", err
	}
	return target, nil
}
