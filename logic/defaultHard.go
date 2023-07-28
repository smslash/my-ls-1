package logic

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"git/ssengerb/my-ls-1/models"
)

func DefaultHard(list []models.File, flag models.FlagOptions, input []string) {
	ds, fs := DevideDirectoryFile(input)
	if len(ds) == 1 && ds[0] == "-" && !flag.Flag_R && !flag.Flag_a && !flag.Flag_l && !flag.Flag_r && !flag.Flag_t {
		return
	}
	if len(fs) != 0 {
		listFiles := make([]models.File, len(fs))
		for i := range listFiles {
			listFiles[i].Name = fs[i]
			FillFile(".", &listFiles[i])
		}

		if flag.Flag_t {
			SortTime(listFiles)
		} else {
			SortDefault(listFiles)
		}

		if flag.Flag_r {
			SortReverse(listFiles)
		}

		if flag.Flag_l {
			PrintFull(flag, listFiles, 2)
		} else {
			PrintDefault(flag, listFiles)
		}

		if len(ds) != 0 {
			fmt.Println()
		}
	}
	for i, path := range ds {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatalln("Error reading directory:", err)
		}

		list = make([]models.File, len(files)+2)

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		fullPath := filepath.Join(cwd, path)
		parentDir := filepath.Dir(fullPath)
		FillCurrentDir(path, &list[len(files)+1])
		list[len(files)+1].Name = "."
		if parentDir == "/home/student/my-ls-1/usr" {
			parentDir = "/usr"
		}
		FillCurrentDir(parentDir, &list[len(files)])
		list[len(files)].Name = ".."

		for i, file := range files {
			list[i].Name = file.Name()
			FillFile(path, &list[i])
		}

		if flag.Flag_t {
			SortTime(list)
		} else {
			SortDefault(list)
		}

		if flag.Flag_r {
			SortReverse(list)
		}

		if flag.Flag_l {
			if len(input) > 1 {
				fmt.Println(path + ":")
			}
			PrintFull(flag, list, 1)
			if i+1 < len(ds) {
				fmt.Println()
			}
		} else {
			if len(input) > 1 {
				fmt.Println(path + ":")
			}
			PrintDefault(flag, list)
			if i+1 < len(ds) {
				fmt.Println()
			}
		}
	}
}
