package file

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory:%v", err)
	}
	defer os.Remove(tempDir)

	tempFile1, err := os.CreateTemp(tempDir, "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile1.Name())

	content1 := "name,age\ntaro,12\nrika,3"
	if _, err := tempFile1.WriteString(content1); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile1.Close()

	tempFile2, err := os.CreateTemp(tempDir, "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile2.Name())

	content2 := "name,age\nmayumi,22\nmiyu,33\nkenji,11"
	if _, err := tempFile2.WriteString(content2); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile2.Close()

	contents, err := Read(tempDir, 1)
	if err != nil {
		t.Errorf("Read returned an error: %v", err)
	}

	if len(contents) != 6 {
		t.Errorf("Expected rows 6, but got %d", len(contents))
	}
}
