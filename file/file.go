package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Read(folderPath string) ([]string, error) {
	var contents []string

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
		return contents, err
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(folderPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				return contents, err
			}
			contents = append(contents, string(content))
		}
	}

	return contents, nil
}

func Make(filePath string, data string) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return err
	}

	return nil
}