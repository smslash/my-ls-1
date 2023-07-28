package models

import (
	"time"
)

type FlagOptions struct {
	Flag_l bool
	Flag_R bool
	Flag_a bool
	Flag_r bool
	Flag_t bool
}

type File struct {
	Permissions   string
	Links         uint64
	Owner         string
	Group         string
	Size          int64
	Total         int64
	Time          time.Time
	Name          string
	IsHidden      bool
	IsDir         bool
	IsFile        bool
	IsSymlink     bool
	SymlinkTarget string
}

type PrettyFormat struct {
	MaxOwnerName int
	MaxGroupName int
	MaxSize      int
}
