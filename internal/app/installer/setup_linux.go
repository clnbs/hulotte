// +build linux

package installer

import "github.com/clnbs/hulotte/internal/app/helper"

const serviceFile = `[Unit]
Description=Hulotte agent

[Service]
ExecStart=/usr/bin/hulotte
ExecReload=/bin/kill -s HUP $MAINPID
TimeoutSec=0
RestartSec=2
Restart=always
`

func SetAutoStart() error {
	return helper.WriteFile([]byte(serviceFile), "/usr/lib/systemd/system/hullote.service")
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
