package KlasorDosyaIslemleri

import (
	"fmt"
	"os"
)

func WriteToFile(data string, filePath string) error {
	// Dosyayı açar veya oluşturur
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open or create file: %v", err)
	}
	defer file.Close()

	// Dosyaya .env yazma işlemi
	if _, err := file.WriteString(fmt.Sprintf("%s \n", data)); err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}

func CreateFolder(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Klasör yoksa oluşturur
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			return fmt.Errorf("failed to create folder: %v", err)
		}
	}

	return nil
}

func IsFileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func IsFolderExists(folderName string) bool {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
