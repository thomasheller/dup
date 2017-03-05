package main

import (
	"log"
	"os"
	"path/filepath"
)

// Walker walks directory hierarchies; returns filenames.
type Walker interface {
	Walk(root string) []string
}

// FilePathWalker walks directory hierarchies using Walk from
// path/filepath.
type FilePathWalker struct{}

// Walk returns all filenames (recursive) starting at root, ignoring
// directories.
func (w FilePathWalker) Walk(root string) []string {
	fileNames := make([]string, 0)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error walking file path %s: %v", path, err)
		}

		if info.IsDir() {
			return nil
		}

		fileNames = append(fileNames, path)

		return nil
	})

	return fileNames
}
