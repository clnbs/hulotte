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

func GetHulottePath() string {
	return "/System/Library/LaunchAgents/hulotte"
}

func RegisterHulotte() error {
	// does nothing
	// we don't need to register a program on macos systems
	return nil
}
