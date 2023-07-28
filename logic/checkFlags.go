package logic

import (
	"fmt"
	"os"

	"git/ssengerb/my-ls-1/models"
)

func CheckFlagsAndInput(s []string) (models.FlagOptions, []string) {
	input := []string{}
	var flag models.FlagOptions
	for i := 0; i < len(s); i++ {
		if s[i][0] == '-' && len(s[i]) > 1 {
			for j := 1; j < len(s[i]); j++ {
				switch s[i][j] {
				case 'l':
					flag.Flag_l = true
				case 'R':
					flag.Flag_R = true
				case 'a':
					flag.Flag_a = true
				case 't':
					flag.Flag_t = true
				case 'r':
					flag.Flag_r = true
				default:
					fmt.Print("-l\n-R\n-a\n-r\n-t\nUse only this flags!\n")
					os.Exit(1)
				}
			}
		} else {
			input = append(input, s[i])
		}
	}
	return flag, input
}
