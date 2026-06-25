package exif

import (
	"errors"
	"os/exec"
)

func CheckInstalled() error {

	path, err := exec.LookPath("exiftool")

	if err != nil {
		return errors.New(
			"ExifTool was not found in PATH.\n" +
				"Download it and add it to your PATH environment variable.",
		)
	}

	cmd := exec.Command(path, "-ver")

	if err := cmd.Run(); err != nil {
		return errors.New("ExifTool exists but cannot be executed")
	}

	return nil
}
