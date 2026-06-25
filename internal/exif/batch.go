package exif

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"os/exec"
	"path/filepath"
	"renamer/internal/model"
	"time"
)

func ReadVideos(paths []string) ([]model.Video, error) {

	if len(paths) == 0 {
		return nil, errors.New("no files provided")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	args := []string{
		"-j",
		"-MediaCreateDate",
		"-CreateDate",
		"-TrackCreateDate",
		"-FileModifyDate",
	}

	args = append(args, paths...)

	cmd := exec.CommandContext(
		ctx,
		"exiftool",
		args...,
	)

	output, err := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("ExifTool timed out")
	}

	if err != nil {
		return nil, err
	}

	var result []exifData

	if err := json.Unmarshal(output, &result); err != nil {

		return nil, fmt.Errorf(
			"failed to parse ExifTool JSON (debug-exif.json created): %w",
			err,
		)
	}

	var videos []model.Video

	for _, item := range result {

		timestamp, err := selectTimestamp(item)
		if err != nil {
			continue
		}

		videos = append(videos, model.Video{
			Path:      item.SourceFile,
			Name:      filepath.Base(item.SourceFile),
			Extension: filepath.Ext(item.SourceFile),
			Time:      timestamp,
		})
	}

	if len(videos) == 0 {
		return nil, errors.New("no valid videos found")
	}

	return videos, nil
}
