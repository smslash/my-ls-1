package logic

import (
	"fmt"
	"log"
	"os"
)

func IsExistFileOrDir(input []string) {
	dir, err := os.Open(".")
	if err != nil {
		log.Fatalln("Error opening directory:", err)
	}
	defer dir.Close()

	entries, err := dir.ReadDir(0)
	if err != nil {
		log.Fatalln("Error reading directory:", err)
	}

	for _, v := range input {
		find := false
		for _, entry := range entries {
			if v == entry.Name() {
				find = true
			}
		}

		if !find {
			fmt.Println("my-ls: cannot access '" + v + "': No such file or directory")
			os.Exit(1)
		}
	}
}
