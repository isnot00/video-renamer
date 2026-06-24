package exif

import (
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"path/filepath"
	"time"

	"renamer/internal/model"
)

func mapToVideos(
	data []exifData,
	pathMap map[string]string,
) ([]model.Video, error) {

	var videos []model.Video

	for _, item := range data {

		t, err := selectTimestamp(item)
		if err != nil {
			continue
		}

		realPath, ok := pathMap[filepath.Base(item.SourceFile)]

		if !ok {
			continue
		}

		videos = append(videos, model.Video{
			Path:      realPath,
			Name:      filepath.Base(realPath),
			Extension: filepath.Ext(realPath),
			Time:      t,
		})
	}

	if len(videos) == 0 {
		return nil, errors.New("no valid videos found")
	}

	return videos, nil
}

func ReadVideos(paths []string) ([]model.Video, error) {

	if len(paths) == 0 {
		return nil, errors.New("no files provided")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		20*time.Second,
	)
	defer cancel()

	pathMap := make(map[string]string)

	var absolutePaths []string

	for _, p := range paths {

		abs, err := filepath.Abs(p)
		if err != nil {
			return nil, err
		}

		pathMap[filepath.Base(abs)] = abs
		absolutePaths = append(
			absolutePaths,
			abs,
		)
	}

	args := []string{
		"-j",
		"-MediaCreateDate",
		"-CreateDate",
		"-TrackCreateDate",
		"-FileModifyDate",
	}

	args = append(args, absolutePaths...)

	cmd := exec.CommandContext(
		ctx,
		"exiftool",
		args...,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {

		if len(output) > 0 {
			return nil, errors.New(string(output))
		}

		return nil, err
	}

	var result []exifData

	if err := json.Unmarshal(
		output,
		&result,
	); err != nil {
		return nil, err
	}

	return mapToVideos(
		result,
		pathMap,
	)
}
