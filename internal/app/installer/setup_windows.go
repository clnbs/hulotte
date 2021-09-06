package installer

import (
	"fmt"
	"os"

	"github.com/clnbs/hulotte/internal/app/helper"

	"golang.org/x/sys/windows/registry"
)

func SetAutoStart() error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	autoStartFolder := "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu \\Programs\\Startup"
	rootAutoStartFolder := fmt.Sprintf("%s%s", userHome, autoStartFolder)
	hulotteInstallationPath := "C:\\Program Files\\hulotte\\hulotte.exe"

	return os.Symlink(hulotteInstallationPath, rootAutoStartFolder+"\\hulotte.exe")
}

func WriteHulotte(data []byte) error {
	err := helper.CreateDir("C:\\Program Files\\hulotte")
	if err != nil {
		return err
	}
	return helper.WriteBinary(data, "C:\\Program Files\\hulotte\\hulotte.exe")
}

func GetHulottePath() string {
	return "C:\\Program Files\\hulotte\\hulotte.exe"
}

func RegisterHulotte() error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Program Files\hulotte\hulotte.exe`, registry.WRITE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	if err = k.SetStringValue("Path", "C:\\Program Files\\hulotte\\hulotte.exe"); err != nil {
		return err
	}

	_, err = registry.OpenKey(registry.CLASSES_ROOT, `Applications\hulotte.exe`, registry.WRITE)
	if err != nil {
		return err
	}

	return nil
}
