package rootpath

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

const ModFile = "go.mod"

// Dir returns the root directory
func Dir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrapf(err, "error while getting the current working directory")
	}

	for {
		has, err := hasTarget(wd, ModFile)
		if err != nil {
			return "", errors.Wrapf(err, "error while checking target %q", wd)
		}
		if has {
			return wd, nil
		}
		wd2 := filepath.Dir(wd)

		if wd == wd2 {
			return "", fmt.Errorf("can't find the root directory containing %q", ModFile)
		}

		wd = wd2
	}
}

// MustDir returns the root directory or panics
func MustDir() string {
	dir, err := Dir()
	if err != nil {
		panic(err)
	}

	return dir
}

// MustChdir changes the current working directory to the root directory
func MustChdir() {
	if err := os.Chdir(MustDir()); err != nil {
		panic(err)
	}
}

func hasTarget(source, target string) (bool, error) {
	files, err := os.ReadDir(source)
	if err != nil {
		return false, errors.Wrapf(err, "error while reading a dir %q", source)
	}

	for _, file := range files {
		if file.Name() == target {
			return true, nil
		}
	}

	return false, nil
}
