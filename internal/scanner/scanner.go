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

	var files []string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(d.Name()))

		if _, ok := SupportedExtensions[ext]; !ok {
			return nil
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
