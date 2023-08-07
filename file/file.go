package file

import (
	"os"
	"path/filepath"
)

func CheckFile(file string) (bool, error) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, err
}

func ReadFileCont(file string) ([]byte, error) {
	checkfile, errorcheck := CheckFile(file)
	if !checkfile || errorcheck != nil {
		return nil, errorcheck
	}
	cont, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return cont, nil
}
func ReWriteFile(file, text string) error {
	newfile, err := os.Create(file)
	if err != nil {
		return err
	}
	_, err = newfile.WriteString(text)
	if err != nil {
		return err
	}
	err = newfile.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetBinaryDir() (string, error) {
	binpath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(binpath), nil
}
