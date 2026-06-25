package exif

import (
	"errors"
	"time"
)

type exifData struct {
	SourceFile      string `json:"SourceFile"`
	MediaCreateDate string `json:"MediaCreateDate"`
	CreateDate      string `json:"CreateDate"`
	TrackCreateDate string `json:"TrackCreateDate"`
	FileModifyDate  string `json:"FileModifyDate"`
}

func parseTime(value string) (time.Time, error) {

	if value == "" {
		return time.Time{}, errors.New("empty timestamp")
	}

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

	return time.Time{}, errors.New("unsupported timestamp: " + value)
}

func selectTimestamp(data exifData) (time.Time, error) {

	fields := []string{
		data.MediaCreateDate,
		data.CreateDate,
		data.TrackCreateDate,
		data.FileModifyDate,
	}

	for _, field := range fields {

		if field == "" {
			continue
		}

		t, err := parseTime(field)

		if err == nil {
			return t, nil
		}

	}

	return time.Time{}, errors.New("no usable timestamp found")
}
