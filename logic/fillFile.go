package logic

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"

	"git/ssengerb/my-ls-1/models"
)

var Format models.PrettyFormat

func FillFile(path string, file *models.File) {
	file.Permissions = findPermissions(path, *file)
	file.Links = findLinks(path, *file)
	file.Owner = findOwner(path, *file)
	if len(file.Owner) > Format.MaxOwnerName {
		Format.MaxOwnerName = len(file.Owner)
	}
	file.Group = findGroup(path, *file)
	if len(file.Group) > Format.MaxGroupName {
		Format.MaxGroupName = len(file.Group)
	}
	file.Size = findSize(path, *file)
	if len(strconv.FormatInt(file.Size, 10)) > Format.MaxSize {
		Format.MaxSize = len(strconv.FormatInt(file.Size, 10))
	}
	file.Total = int64(math.Round(float64(file.Size) / 1000))
	file.Time = findTime(path, *file)
	file.IsHidden = isHidden(path, *file)
	file.IsDir = isDir(path, *file)
	file.IsFile = isFile(path, *file)
	file.IsSymlink = isSymlink(path, *file)
	if file.IsSymlink {
		file.SymlinkTarget = findSymlinkTarget(path, *file)
	}
}

func findPermissions(path string, file models.File) string {
	info, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.ToLower(info.Mode().String())
}

func findLinks(path string, file models.File) uint64 {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	sysStat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalln("conversion to *syscall.Stat_t failed")
	}

	return uint64(sysStat.Nlink)
}

func findOwner(path string, file models.File) string {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalln("Failed to get owner information")
	}

	ownerUser, err := user.LookupId(fmt.Sprint(stat.Uid))
	if err != nil {
		log.Fatalln(err)
	}

	return ownerUser.Username
}

func findGroup(path string, file models.File) string {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalln("Failed to get group information")
	}

	group, err := user.LookupGroupId(fmt.Sprint(stat.Gid))
	if err != nil {
		log.Fatalln(err)
	}

	return group.Name
}

func findSize(path string, file models.File) int64 {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.Size()
}

func findTime(path string, file models.File) time.Time {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.ModTime()
}

func isHidden(path string, file models.File) bool {
	base := getBaseName(path + "/" + file.Name)
	return len(base) > 0 && base[0] == '.'
}

func getBaseName(filePath string) string {
	segments := strings.Split(filePath, string(os.PathSeparator))
	return segments[len(segments)-1]
}

func isDir(path string, file models.File) bool {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.IsDir()
}

func isFile(path string, file models.File) bool {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return !fileInfo.IsDir() && fileInfo.Mode().IsRegular()
}

func isSymlink(path string, file models.File) bool {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.Mode()&os.ModeSymlink != 0
}

func findSymlinkTarget(path string, file models.File) string {
	targetPath, err := os.Readlink(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return targetPath
}
