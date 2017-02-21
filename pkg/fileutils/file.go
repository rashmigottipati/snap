package fileutils

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"

	log "github.com/Sirupsen/logrus"
)

// TODO add doc string
func WriteFile(fileName, filePath string, b []byte) (string, error) {
	// Create temporary directory
	dir, err := ioutil.TempDir(filePath, "snap-plugin-")
	if err != nil {
		return "", err
	}

	f, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		return "", err
	}
	// Close before load
	defer f.Close()

	n, err := f.Write(b)
	log.Debugf("wrote %v to %v", n, f.Name())
	if err != nil {
		return "", err
	}
	if runtime.GOOS != "windows" {
		err = f.Chmod(0700)
		if err != nil {
			return "", err
		}
	}
	return f.Name(), nil
}
