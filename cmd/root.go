package cmd

import (
	"fmt"
	"log"
	"merge-txt/file"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var path string
var row int

var rootCmd = &cobra.Command{
	Use:   "merge-txt",
	Short: "Combine text files in any folder",
	Long:  "Combine text files in any folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		contents, err := file.Read(path, row)
		if err != nil {
			log.Fatalln(err)
		}

		filePath := filepath.Join(path, "output.txt")
		data := strings.Join(contents, "\n")

		err = file.Make(filePath, data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Combined file saved to %s\n", filePath)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// フラグの値を変数にバインド
	rootCmd.Flags().StringVar(&path, "path", "", "File Location")
	rootCmd.Flags().IntVar(&row, "row", 1, "Header Lines")

	// 必須のフラグに指定
	rootCmd.MarkFlagRequired("path")
}
