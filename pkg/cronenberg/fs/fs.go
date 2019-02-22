// Package fs provides functions for working with the file system as well as
// services that specifically deal with the file system.
package fs

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

// Root is an Afero Filesystem object used as a proxy to the real file system.
var Root = afero.NewOsFs()

// CreateDir takes a path and a file mode, and it creates the path on Root with
// the given file mode. If there are any issues, an error is returned.
// Otherwise, nil is returned.
var CreateDir = func(path string, mode os.FileMode) error {
	if !FileExists(path) {
		err := Root.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

// FileExists takes a file path and returns a bool. If the file exists on Root,
// the result is true. Otherwise, it is false.
var FileExists = func(path string) bool {
	_, err := Root.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// DirectoryExists takes a directory path and returns a bool. If the directory
// exists on Root, the result is true. Otherwise, it is false.
var DirectoryExists = func(path string) bool {
	if !FileExists(path) {
		return false
	}

	if !IsDir(path) {
		return false
	}

	return true
}

// IsDir takes a path and returns a bool. If the path is a directory, the result
// is true. Otherwise, it is false.
var IsDir = func(path string) bool {
	info, err := Root.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

// ReadFile takes a file path, reads that file, and returns a byte array of the
// file's contents and an error. If there are issues, the error is non-nil.
// Otherwise, the error is nil.
func ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(Root, path)
}

// Basename takes a file path and returns the actual filename.
func Basename(path string) string {
	return filepath.Base(path)
}

// Stat takes a path and returns both a file info object and an error. If there
// are issues along the way, the error is non-nil. Otherwise, the error is nil.
func Stat(path string) (os.FileInfo, error) {
	return Root.Stat(path)
}

// Executable takes a file path and returns a bool. If the file's permissions
// have the executable bit enabled, the result is true. Otherwise, it is false.
func Executable(path string) bool {
	info, _ := Stat(path)

	return (info.Mode()&0100 == 0100)
}

/*
Copyright 2019 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
