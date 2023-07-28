package logic

import (
	"fmt"

	"git/ssengerb/my-ls-1/models"
)

func PrintDefault(flag models.FlagOptions, files []models.File) {
	if flag.Flag_a {
		for i := 0; i < len(files); i++ {
			if files[i].IsDir {
				fmt.Print("\033[34m" + files[i].Name + "\033[0m")
			} else if files[i].IsSymlink {
				fmt.Print("\033[96m" + files[i].Name + "\033[0m")
			} else {
				fmt.Print(files[i].Name)
			}
			if i+1 < len(files) {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	} else {
		for i := 0; i < len(files); i++ {
			if !files[i].IsHidden {
				if files[i].IsDir {
					fmt.Print("\033[34m" + files[i].Name + "\033[0m")
				} else if files[i].IsSymlink {
					fmt.Print("\033[96m" + files[i].Name + "\033[0m")
				} else {
					fmt.Print(files[i].Name)
				}
				if i+1 < len(files) {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
}
