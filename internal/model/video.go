package model

import "time"

type Video struct {
	Path       string
	Name       string
	Extension  string
	Time       time.Time
	TargetName string
}
