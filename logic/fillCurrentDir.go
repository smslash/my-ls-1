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

func FillCurrentDir(path string, file *models.File) {
	file.Permissions = fillPermissions(path, *file)
	file.Links = fillLinks(path, *file)
	file.Owner = fillOwner(path, *file)
	if len(file.Owner) > Format.MaxOwnerName {
		Format.MaxOwnerName = len(file.Owner)
	}
	file.Group = fillGroup(path, *file)
	if len(file.Group) > Format.MaxGroupName {
		Format.MaxGroupName = len(file.Group)
	}
	file.Size = fillSize(path, *file)
	if len(strconv.FormatInt(file.Size, 10)) > Format.MaxSize {
		Format.MaxSize = len(strconv.FormatInt(file.Size, 10))
	}
	file.Total = int64(math.Round(float64(file.Size) / 1000))
	file.Time = fillTime(path, *file)
	file.IsHidden = true
	file.IsDir = iDir(path, *file)
	file.IsFile = iFile(path, *file)
	file.IsSymlink = iSymlink(path, *file)
	if file.IsSymlink {
		file.SymlinkTarget = fillSymlinkTarget(path, *file)
	}
}

func fillPermissions(path string, file models.File) string {
	info, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.ToLower(info.Mode().String())
}

func fillLinks(path string, file models.File) uint64 {
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

func fillOwner(path string, file models.File) string {
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

func fillGroup(path string, file models.File) string {
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

func fillSize(path string, file models.File) int64 {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.Size()
}

func fillTime(path string, file models.File) time.Time {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.ModTime()
}

func iDir(path string, file models.File) bool {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.IsDir()
}

func iFile(path string, file models.File) bool {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return !fileInfo.IsDir() && fileInfo.Mode().IsRegular()
}

func iSymlink(path string, file models.File) bool {
	fileInfo, err := os.Lstat(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return fileInfo.Mode()&os.ModeSymlink != 0
}

func fillSymlinkTarget(path string, file models.File) string {
	targetPath, err := os.Readlink(path + "/" + file.Name)
	if err != nil {
		log.Fatalln(err)
	}

	return targetPath
}
