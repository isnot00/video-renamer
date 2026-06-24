package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"renamer/internal/model"
	"sort"
)

func SortVideos(videos []model.Video) {
	sort.Slice(videos, func(i, j int) bool {
		return videos[i].Time.Before(videos[j].Time)
	})
}

func GenerateNames(videos []model.Video) {
	for i := range videos {
		videos[i].TargetName = fmt.Sprintf(
			"%04d%s",
			i+1,
			videos[i].Extension,
		)
	}
}

func ValidateTargets(videos []model.Video) error {

	targets := make(map[string]struct{})

	for _, video := range videos {

		finalPath := filepath.Join(
			filepath.Dir(video.Path),
			video.TargetName,
		)

		if _, exists := targets[finalPath]; exists {
			return fmt.Errorf(
				"duplicate target generated: %s",
				finalPath,
			)
		}

		targets[finalPath] = struct{}{}
	}

	return nil
}

func PrintPreview(videos []model.Video) {

	fmt.Println()
	fmt.Printf(
		"%-30s %-20s\n",
		"OLD FILE",
		"NEW FILE",
	)

	fmt.Println(
		"────────────────────────────────────────────",
	)

	for _, v := range videos {

		fmt.Printf(
			"%-30s %-20s\n",
			v.Name,
			v.TargetName,
		)
	}

	fmt.Println()
}

func Rename(videos []model.Video, dryRun bool) error {

	for i := range videos {

		dir := filepath.Dir(videos[i].Path)

		tmpPath := filepath.Join(
			dir,
			fmt.Sprintf(
				"__video_renamer_tmp_%d_%s",
				i,
				videos[i].Name,
			),
		)

		if dryRun {
			fmt.Printf(
				"[DRY RUN] %s -> %s\n",
				videos[i].Name,
				videos[i].TargetName,
			)
			continue
		}

		if err := os.Rename(videos[i].Path, tmpPath); err != nil {
			return err
		}

		videos[i].Path = tmpPath
	}

	if dryRun {
		return nil
	}

	for i := range videos {

		dir := filepath.Dir(videos[i].Path)

		finalPath := filepath.Join(
			dir,
			videos[i].TargetName,
		)

		if err := os.Rename(
			videos[i].Path,
			finalPath,
		); err != nil {
			return err
		}
	}

	return nil
}
