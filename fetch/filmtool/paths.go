package main

import (
	"errors"
	"os"
	"path/filepath"
)

// findRoot walks up from the working directory to the repository root,
// identified by eleventy.config.js, so the tool can be run from anywhere
// inside the repo.
func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "eleventy.config.js")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("not inside the blog repository (eleventy.config.js not found)")
		}
		dir = parent
	}
}

func filmDataPath(root, slug string) string {
	return filepath.Join(root, "_data", "films", slug+".json")
}

func imagePath(root, kind, name string) string {
	return filepath.Join(root, "content", "bill", "films", kind, name)
}
