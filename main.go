package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func readFile(folderPath string) ([]string, error) {
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

func makeFile(filePath string, data string) error {
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

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}

	exeDir := filepath.Dir(exePath)
	folderPath := filepath.Join(exeDir, "data")

	contents, err := readFile(folderPath)
	if err != nil {
		log.Fatalln(err)
	}

	filePath := filepath.Join(exeDir, "output.txt")
	data := strings.Join(contents, "\n")

	err = makeFile(filePath, data)
	if err != nil {
		log.Fatalln(err)
	}
}
