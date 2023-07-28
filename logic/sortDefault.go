package logic

import (
	"strings"
	"unicode"

	"git/ssengerb/my-ls-1/models"
)

func SortDefault(files []models.File) {
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(files)-i-1; j++ {
			if files[j].IsHidden && !files[j+1].IsHidden {
				continue
			}
			if !files[j].IsHidden && files[j+1].IsHidden {
				files[j], files[j+1] = files[j+1], files[j]
			} else if !lsSort(strings.ToLower(files[j].Name), strings.ToLower(files[j+1].Name)) {
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
}

func lsSort(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	for i := 0; i < len(r1) && i < len(r2); i++ {
		isRune1Letter := unicode.IsLetter(r1[i]) || unicode.IsNumber(r1[i])
		isRune2Letter := unicode.IsLetter(r2[i]) || unicode.IsNumber(r2[i])

		if isRune1Letter != isRune2Letter {
			return isRune2Letter
		}

		if r1[i] != r2[i] {
			return r1[i] < r2[i]
		}
	}
	return len(r1) < len(r2)
}
