package dkr

import (
	"os"
	"path"
)

type dkrHome struct {
	Root string
}

func NewDcmHome() (*dkrHome, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	root := path.Join(userHome, ".dkr")
	return NewDcmHomeWihRoot(root)
}

func NewDcmHomeWihRoot(root string) (*dkrHome, error) {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		err = os.Mkdir(root, 0700)
		if err != nil {
			return nil, err
		}
	}

	return &dkrHome{Root: root}, nil
}
