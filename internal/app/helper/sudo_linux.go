// +build linux
package helper

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func SUDOWriteBinary(data []byte, path string) error {
	tmpBinaryPath, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	err = WriteBinary(data, tmpBinaryPath+"/tmp_hulotte_bin")
	if err != nil {
		return err
	}

	is, err := DoesFileExsist(tmpBinaryPath + "/tmp_hulotte_bin")
	if err != nil {
		return err
	}
	if !is {
		return errors.New("temporary binary file not found")
	}

	stringifyCommand := fmt.Sprintf("sudo mv %s %s", tmpBinaryPath+"/tmp_hulotte_bin", path)
	cmd := exec.Command("/bin/sh", "-c", stringifyCommand)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func SUDOWriteFile(data []byte, path string) error {
	tmpFilePath, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	err = WriteFile(data, tmpFilePath+"/tmp_hulotte_file")
	if err != nil {
		return err
	}

	is, err := DoesFileExsist(tmpFilePath + "/tmp_hulotte_file")
	if err != nil {
		return err
	}
	if !is {
		return errors.New("temporary file not found")
	}

	stringifyCommand := fmt.Sprintf("sudo mv %s %s", tmpFilePath+"/tmp_hulotte_file", path)
	cmd := exec.Command("/bin/sh", "-c", stringifyCommand)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
