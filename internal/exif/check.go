package exif

import "os/exec"

func CheckInstalled() error {
	_, err := exec.LookPath("exiftool")
	return err
}
