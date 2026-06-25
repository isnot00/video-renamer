package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"renamer/internal/config"
	"renamer/internal/exif"
	"renamer/internal/renamer"
	"renamer/internal/scanner"
	"renamer/internal/ui"
)

func waitExit() {
	fmt.Println()
	fmt.Print("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func main() {

	ui.PrintBanner()

	dir := flag.String(
		"dir",
		"",
		"directory containing videos",
	)

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	if *dir == "" {
		fmt.Print("Video folder: ")

		text, _ := reader.ReadString('\n')

		*dir = strings.TrimSpace(text)
	}

	cfg := config.Config{
		Directory: *dir,
	}

	if cfg.Directory == "" {
		ui.Error("directory is required")
		waitExit()
		return
	}

	if err := exif.CheckInstalled(); err != nil {
		ui.Error("ExifTool not found.")
		waitExit()
		return
	}

	ui.Info("Scanning directory...")

	paths, err := scanner.ScanDirectory(cfg.Directory)

	if err != nil {
		ui.Error(err.Error())
		waitExit()
		return
	}

	if len(paths) == 0 {
		ui.Warning("No video files found.")
		waitExit()
		return
	}

	ui.Success(
		fmt.Sprintf(
			"Found %d video files",
			len(paths),
		),
	)

	ui.Info("Reading metadata...")

	videos, err := exif.ReadVideos(paths)

	if err != nil {
		ui.Error(err.Error())
		waitExit()
		return
	}

	ui.Success(
		fmt.Sprintf(
			"Read metadata for %d videos",
			len(videos),
		),
	)

	ui.Info("Sorting videos...")
	renamer.SortVideos(videos)

	ui.Info("Generating file names...")
	renamer.GenerateNames(videos)

	if err := renamer.ValidateTargets(videos); err != nil {
		ui.Error(err.Error())
		waitExit()
		return
	}

	renamer.PrintPreview(videos)

	fmt.Print("\nRename these files? [y/N]: ")

	answer, _ := reader.ReadString('\n')

	answer = strings.TrimSpace(strings.ToLower(answer))

	if answer != "y" && answer != "yes" {
		ui.Warning("Operation cancelled.")
		waitExit()
		return
	}

	ui.Info("Renaming files...")

	if err := renamer.Rename(videos); err != nil {
		ui.Error(err.Error())
		waitExit()
		return
	}

	ui.Success("Rename completed successfully.")

	waitExit()
}
