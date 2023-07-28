package logic

import "git/ssengerb/my-ls-1/models"

func SortTime(files []models.File) {
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(files)-i-1; j++ {
			if files[j].Time.Before(files[j+1].Time) {
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
}
