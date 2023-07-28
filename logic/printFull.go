package logic

import (
	"fmt"
	"strconv"

	"git/ssengerb/my-ls-1/models"
)

func PrintFull(flag models.FlagOptions, files []models.File, num int) {
	if flag.Flag_a {
		var total int64
		for i := 0; i < len(files); i++ {
			total += files[i].Total
		}

		if num == 1 {
			fmt.Println("total", total)
		}

		for i := 0; i < len(files); i++ {
			fmt.Print(files[i].Permissions + " ")
			if len(files[i].Permissions) <= 10 {
				fmt.Print(" ")
			}
			if len(strconv.FormatUint(files[i].Links, 10)) == 1 {
				fmt.Print(" ")
			}
			fmt.Print(files[i].Links)
			fmt.Print(" ")
			if len(files[i].Owner) < Format.MaxOwnerName {
				for j := 0; j < Format.MaxOwnerName-len(files[i].Owner); j++ {
					fmt.Print(" ")
				}
			}
			fmt.Print(files[i].Owner + " ")
			fmt.Print(files[i].Group + " ")
			if len(files[i].Group) < Format.MaxGroupName {
				for j := 0; j < Format.MaxGroupName-len(files[i].Group); j++ {
					fmt.Print(" ")
				}
			}
			if len(strconv.FormatInt(files[i].Size, 10)) < Format.MaxSize {
				for j := 0; j < Format.MaxSize-len(strconv.FormatInt(files[i].Size, 10)); j++ {
					fmt.Print(" ")
				}
			}
			fmt.Print(files[i].Size)
			fmt.Print(" " + kzMonth(files[i].Time.Month().String()) + " ")
			if len(strconv.Itoa(files[i].Time.Day())) == 1 {
				fmt.Print(" ")
			}
			fmt.Print(strconv.Itoa(files[i].Time.Day()) + " ")
			fmt.Print(files[i].Time.Format("15:04") + " ")
			if files[i].IsDir {
				fmt.Print("\033[34m" + files[i].Name + "\033[0m")
			} else if files[i].IsSymlink {
				fmt.Print("\033[96m" + files[i].Name + "\033[0m" + " -> " + files[i].SymlinkTarget)
			} else {
				fmt.Print(files[i].Name)
			}
			fmt.Println()
		}
	} else {
		var total int64
		for i := 0; i < len(files); i++ {
			if !files[i].IsHidden {
				total += files[i].Total
			}
		}

		if num == 1 {
			fmt.Println("total", total)
		}

		for i := 0; i < len(files); i++ {
			if !files[i].IsHidden {
				fmt.Print(files[i].Permissions + " ")
				if len(files[i].Permissions) <= 10 {
					fmt.Print(" ")
				}
				if len(strconv.FormatUint(files[i].Links, 10)) == 1 {
					fmt.Print(" ")
				}
				fmt.Print(files[i].Links)
				fmt.Print(" ")
				if len(files[i].Owner) < Format.MaxOwnerName {
					for j := 0; j < Format.MaxOwnerName-len(files[i].Owner); j++ {
						fmt.Print(" ")
					}
				}
				fmt.Print(files[i].Owner + " ")
				fmt.Print(files[i].Group + " ")
				if len(files[i].Group) < Format.MaxGroupName {
					for j := 0; j < Format.MaxGroupName-len(files[i].Group); j++ {
						fmt.Print(" ")
					}
				}
				if len(strconv.FormatInt(files[i].Size, 10)) < Format.MaxSize {
					for j := 0; j < Format.MaxSize-len(strconv.FormatInt(files[i].Size, 10)); j++ {
						fmt.Print(" ")
					}
				}
				fmt.Print(files[i].Size)
				fmt.Print(" " + kzMonth(files[i].Time.Month().String()) + " ")
				if len(strconv.Itoa(files[i].Time.Day())) == 1 {
					fmt.Print(" ")
				}
				fmt.Print(strconv.Itoa(files[i].Time.Day()) + " ")
				fmt.Print(files[i].Time.Format("15:04") + " ")
				if files[i].IsDir {
					fmt.Print("\033[34m" + files[i].Name + "\033[0m")
				} else if files[i].IsSymlink {
					fmt.Print("\033[96m" + files[i].Name + "\033[0m" + " -> " + files[i].SymlinkTarget)
				} else {
					fmt.Print(files[i].Name)
				}
				fmt.Println()
			}
		}
	}
}

func kzMonth(s string) string {
	if s == "January" {
		return "Қаң"
	} else if s == "February" {
		return "Ақп"
	} else if s == "March" {
		return "Нау"
	} else if s == "April" {
		return "Сәу"
	} else if s == "May" {
		return "Мам"
	} else if s == "June" {
		return "Мау"
	} else if s == "July" {
		return "Шiл"
	} else if s == "August" {
		return "Там"
	} else if s == "September" {
		return "Қыр"
	} else if s == "October" {
		return "Қаз"
	} else if s == "November" {
		return "Қар"
	}
	return "Жел"
}
