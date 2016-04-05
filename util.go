package deanio

import (
	"os/user"
)

var homeDir string

func initUtil() error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	homeDir = u.HomeDir + ".deanio/"
	return nil
}
