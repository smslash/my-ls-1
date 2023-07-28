package logic

import (
	"log"
	"os"
)

func DevideDirectoryFile(s []string) ([]string, []string) {
	var ds []string
	var fs []string

	for _, path := range s {
		fileInfo, err := os.Lstat(path)
		if err != nil {
			log.Fatalln("Error:", err)
		}

		if fileInfo.Mode().IsRegular() {
			fs = append(fs, path)
		} else if fileInfo.Mode().IsDir() {
			ds = append(ds, path)
		} else if fileInfo.Mode()&os.ModeSymlink != 0 {
			fs = append(fs, path)
		}
	}

	return ds, fs
}
