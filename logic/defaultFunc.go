package logic

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"git/ssengerb/my-ls-1/models"
)

func DefaultFunc(list []models.File, flag models.FlagOptions, path string) {
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
		PrintFull(flag, list, 1)
	} else {
		PrintDefault(flag, list)
	}
}
