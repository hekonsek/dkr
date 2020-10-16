package dkr

import (
	"os"
	"path"
)

type dkrHome struct {
	root string
}

func (home *dkrHome) Root() string {
	return home.root
}

func (home *dkrHome) Bin() string {
	return path.Join(home.Root(), "bin")
}

func NewDkrHome() (*dkrHome, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	root := path.Join(userHome, ".dkr")
	return NewDkrHomeWihRoot(root)
}

func NewDkrHomeWihRoot(root string) (*dkrHome, error) {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		err = os.Mkdir(root, 0700)
		if err != nil {
			return nil, err
		}
	}

	dkr := &dkrHome{root: root}

	if _, err := os.Stat(dkr.Bin()); os.IsNotExist(err) {
		err = os.Mkdir(dkr.Bin(), 0700)
		if err != nil {
			return nil, err
		}
	}

	return dkr, nil
}
