// +build linux

package installer

import "github.com/clnbs/hulotte/internal/app/helper"

func SetAutoStart() error {
	// TODO thumbs
	return nil
}

func WriteHulotte(data []byte) error {
	return helper.WriteBinary(data, "/usr/bin/hulotte")
}

func GetHullotePath() string {
	return "/usr/bin/hulotte"
}

func RegisterHulotte() error {
	// does nothing
	// we don't need to register a program on linux systems
	return nil
}
