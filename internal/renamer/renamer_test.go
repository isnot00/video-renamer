package renamer

import (
	"renamer/internal/model"
	"testing"
)

func TestGenerateNames(t *testing.T) {

	videos := []model.Video{
		{Extension: ".mp4"},
		{Extension: ".mp4"},
		{Extension: ".mov"},
	}

	GenerateNames(videos)

	expected := []string{
		"0001.mp4",
		"0002.mp4",
		"0003.mov",
	}

	for i := range videos {

		if videos[i].TargetName != expected[i] {

			t.Fatalf(
				"expected %s got %s",
				expected[i],
				videos[i].TargetName,
			)
		}
	}
}
