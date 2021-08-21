package helper

import "os"

func WriteFile(data []byte, path string) error {
	err := os.WriteFile(path, data, 0766)
	if err != nil {
		return err
	}
	return nil
}

func WriteBinary(data []byte, path string) error {
	err := os.WriteFile(path, data, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CreateDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func DoesFileExsist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
