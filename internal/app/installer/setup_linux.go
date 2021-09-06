package installer

import (
	"fmt"
	"os"

	"github.com/clnbs/hulotte/internal/app/helper"
)

const autostartFile = `[Desktop Entry]
Type=Application
Name=Hulotte
Exec=/usr/bin/hulotte
StartupNotify=false
Terminal=false
`

func SetAutoStart() error {
	autostartFilePath, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	autostartFilePath += "/autostart/hulotte.desktop "
	fmt.Println("Writing service file for Linux in", autostartFilePath)

	return helper.SUDOWriteFile([]byte(autostartFile), autostartFilePath)
}

func WriteHulotte(data []byte) error {
	fmt.Println("Writing Hulotte binary for Linux in /usr/bin/hulotte")
	return helper.SUDOWriteBinary(data, "/usr/bin/hulotte")
}

func GetHulottePath() string {
	return "/usr/bin/hulotte"
}

func RegisterHulotte() error {
	// does nothing
	// we don't need to register a program on linux systems
	fmt.Println("We don't need to register Hulotte on Linux systems :-)")
	return nil
}
