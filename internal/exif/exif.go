package exif

import (
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"path/filepath"
	"renamer/internal/model"
	"time"
)

const timeLayout = "2006:01:02 15:04:05"

type exifData struct {
	SourceFile      string `json:"SourceFile"`
	MediaCreateDate string `json:"MediaCreateDate"`
	CreateDate      string `json:"CreateDate"`
	TrackCreateDate string `json:"TrackCreateDate"`
	FileModifyDate  string `json:"FileModifyDate"`
}

func ReadVideo(path string) (model.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"exiftool",
		"-j",
		"-MediaCreateDate",
		"-CreateDate",
		"-TrackCreateDate",
		"-FileModifyDate",
		path,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return model.Video{}, errors.New(string(output))
	}

	var result []exifData

	if err := json.Unmarshal(output, &result); err != nil {
		return model.Video{}, err
	}

	if len(result) == 0 {
		return model.Video{}, errors.New("no metadata returned")
	}

	timestamp, err := selectTimestamp(result[0])
	if err != nil {
		return model.Video{}, err
	}

	return model.Video{
		Path:      path,
		Name:      filepath.Base(path),
		Extension: filepath.Ext(path),
		Time:      timestamp,
	}, nil
}

func parseTime(value string) (time.Time, error) {

	layouts := []string{
		"2006:01:02 15:04:05-07:00",
		"2006:01:02 15:04:05+07:00",
		"2006:01:02 15:04:05",
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, value)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("invalid time format: " + value)
}

func selectTimestamp(data exifData) (time.Time, error) {
	candidates := []string{
		data.MediaCreateDate,
		data.CreateDate,
		data.TrackCreateDate,
		data.FileModifyDate,
	}

	for _, v := range candidates {
		if v == "" {
			continue
		}

		t, err := parseTime(v)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("no valid timestamp found")
}
