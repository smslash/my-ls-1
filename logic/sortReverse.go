package logic

import "git/ssengerb/my-ls-1/models"

func SortReverse(files []models.File) {
	n := len(files)
	for i := 0; i < n/2; i++ {
		files[i], files[n-i-1] = files[n-i-1], files[i]
	}
}
