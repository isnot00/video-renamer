package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"renamer/internal/config"
	"renamer/internal/exif"
	"renamer/internal/logger"
	"renamer/internal/model"
	"renamer/internal/renamer"
	"renamer/internal/scanner"
	"renamer/internal/ui"
)

func main() {
	ui.PrintBanner()
	if len(os.Args) < 2 {
		logger.Info("usage: video-renamer <directory>")
		os.Exit(1)
	}

	dir := flag.String("dir", "", "directory containing videos")
	dryRun := flag.Bool("dry-run", false, "preview changes")
	execute := flag.Bool("execute", false, "perform rename")

	flag.Parse()
	cfg := config.Config{
		Directory: *dir,
		DryRun:    *dryRun,
		Execute:   *execute,
	}
	if cfg.Directory == "" {
		logger.Info("directory is required")
		os.Exit(1)
	}
	if cfg.DryRun == cfg.Execute {
		logger.Info("choose exactly one of --dry-run or --execute")
		os.Exit(1)
	}

	if err := exif.CheckInstalled(); err != nil {
		logger.Info("exiftool not found")
		os.Exit(1)
	}

	paths, err := scanner.ScanDirectory(cfg.Directory)
	if err != nil {
		ui.Error(err.Error())
		return
	}
	ui.Success(
		fmt.Sprintf(
			"Found %d video files",
			len(paths),
		),
	)

	var videos []model.Video

	for i, path := range paths {
		fmt.Printf(
			"[%d/%d] Reading metadata: %s\n",
			i+1,
			len(paths),
			filepath.Base(path),
		)
		video, err := exif.ReadVideo(path)
		if err != nil {
			ui.Error(err.Error())
			continue
		}

		videos = append(videos, video)

	}
	if err := renamer.ValidateTargets(videos); err != nil {
		panic(err)
	}

	if len(videos) == 0 {
		logger.Info("no valid videos found")
		return
	}
	ui.Info("Sorting videos...")
	renamer.SortVideos(videos)
	ui.Info("Generating names...")
	renamer.GenerateNames(videos)
	renamer.PrintPreview(videos)

	if !cfg.DryRun {

		fmt.Print("\nContinue? [y/N]: ")

		var answer string
		fmt.Scanln(&answer)

		if answer != "y" {
			ui.Warning("Operation cancelled")
			return
		}
	}
	err = renamer.Rename(
		videos,
		cfg.DryRun,
	)

	if cfg.DryRun {
		logger.Info("running in dry-run mode")
	}

	if err := renamer.Rename(videos, cfg.DryRun); err != nil {
		panic(err)
	}
	ui.Success("Rename completed")
}
