package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

var SupportedExtensions = map[string]struct{}{
	".mp4": {},
	".mov": {},
	".m4v": {},
	".avi": {},
	".mkv": {},
}

func ScanDirectory(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))

		if _, ok := SupportedExtensions[ext]; !ok {
			continue
		}

		fullPath := filepath.Join(dir, entry.Name())
		files = append(files, fullPath)
	}

	return files, nil
}