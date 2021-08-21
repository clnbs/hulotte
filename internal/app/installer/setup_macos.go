// +build macos

package installer

import "github.com/clnbs/hulotte/internal/app/helper"

func SetAutoStart() error {
	// Nothing to do,
	// Hulotte will auto startup be cause it's install in auto start dir
	return nil
}

func WriteHulotte(data []byte) error {
	return helper.WriteBinary(data, "/System/Library/LaunchAgents/hulotte")
}

func GetHullotePath() string {
	return "/System/Library/LaunchAgents/hulotte"
}

func RegisterHulotte() error {
	// TODO thumbs
	return nil
}

func DoesHulotteExists() (bool, error) {
	// TODO thumb
	return false, nil
}
