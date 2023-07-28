package main

import (
	"fmt"
	"os"

	"git/ssengerb/my-ls-1/logic"
	"git/ssengerb/my-ls-1/models"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "my-ls" {
		fmt.Print("Usage: go run . my-ls [FLAGS]\n\nEX: go run . my-ls -l\n")
	}

	flag, input := logic.CheckFlagsAndInput(os.Args[2:])
	var list []models.File

	if len(input) == 0 {
		if flag.Flag_R {
			logic.ReverseFunc(list, flag, ".")
		} else {
			logic.DefaultFunc(list, flag, ".")
		}
	} else {
		if flag.Flag_R {
			ds, _ := logic.DevideDirectoryFile(input)
			if len(ds) == 1 && ds[0] == "-" && !flag.Flag_R && !flag.Flag_a && !flag.Flag_l && !flag.Flag_r && !flag.Flag_t {
				return
			}
			for i, dir := range ds {
				logic.ReverseFunc(list, flag, dir)
				if i+1 < len(ds) {
					fmt.Println()
				}
			}
		} else {
			logic.DefaultHard(list, flag, input)
		}
	}
}
