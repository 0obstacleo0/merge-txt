package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Read(folderPath string, headerRows int) ([]string, error) {
	var contents []string

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
		return contents, err
	}

	for i, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(folderPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				return contents, err
			}

			txt := strings.Split(string(content), "\n")
			if i > 0 {
				txt = txt[headerRows:]
			}
			contents = append(contents, txt...)
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
