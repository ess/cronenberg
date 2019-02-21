package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

var Root = afero.NewOsFs()

var CreateDir = func(path string, mode os.FileMode) error {
	if !FileExists(path) {
		err := Root.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

var FileExists = func(path string) bool {
	_, err := Root.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

var DirectoryExists = func(path string) bool {
	if !FileExists(path) {
		return false
	}

	if !IsDir(path) {
		return false
	}

	return true
}

var IsDir = func(path string) bool {
	info, err := Root.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func Walk(path string, walkFunc filepath.WalkFunc) error {
	return afero.Walk(Root, path, walkFunc)
}

func Copy(path string, targetPath string, mode os.FileMode) error {
	infile, err := Root.Open(path)
	if err != nil {
		return err
	}
	defer infile.Close()

	if FileExists(targetPath) {
		rmerr := Root.Remove(targetPath)
		if rmerr != nil {
			return rmerr
		}
	}

	outfile, err := Root.Create(targetPath)
	if err != nil {
		return err
	}
	defer func() {
		cerr := outfile.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(outfile, infile)
	if err != nil {
		return err
	}

	err = outfile.Sync()
	return err
}

func DirCopy(source string, target string) error {
	if !FileExists(source) {
		return fmt.Errorf("%s does not exist", source)
	}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		targetPath := target + strings.TrimPrefix(path, source)

		if info.IsDir() {
			return CreateDir(targetPath, info.Mode())
		} else {
			return Copy(path, targetPath, info.Mode())
		}
	}

	return Walk(source, walkFunc)
}

func ReadDir(path string) ([]os.FileInfo, error) {
	return afero.ReadDir(Root, path)
}

func ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(Root, path)
}

func Basename(path string) string {
	return filepath.Base(path)
}

func Stat(path string) (os.FileInfo, error) {
	return Root.Stat(path)
}

func Executable(path string) bool {
	info, _ := Stat(path)

	return (info.Mode()&0100 == 0100)
}
