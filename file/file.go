package file

import (
	"os"
	"path/filepath"
)

// Returns the FileSize and a error
func FileSize(input string) (int64, error) { // filesize,error
	if check, err := CheckFile(input); !check || err != nil {
		return 0, err
	}
	var size int64
	file, err := os.Stat(input)
	if err != nil {
		return 0, err
	}
	if !file.IsDir() {
		return file.Size(), nil
	} else {
		err = filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				size += info.Size()
			}
			return nil
		})
		if err != nil {
			return 0, err
		}
		return size, nil
	}
}

// Returns a true value if file or folder exists, else returns false
func CheckFile(file string) (bool, error) {
	_, err := os.Stat(file) // Get the stat of the file
	if os.IsNotExist(err) { // Check if not exist
		return false, err
	}
	return true, err
}

// Read the file content and return it in bytes
func ReadFileCont(file string) ([]byte, error) { // Return byte data and error
	checkfile, errorcheck := CheckFile(file) // Check if the file exist
	if !checkfile || errorcheck != nil {
		return nil, errorcheck
	}
	cont, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return cont, nil
}

// Rewrite the file as string type
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

// Returns the running binary diretory, and a error
func GetBinaryDir() (string, error) {
	binpath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(binpath), nil
}
