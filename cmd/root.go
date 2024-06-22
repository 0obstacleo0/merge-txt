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
	Short: "任意のフォルダ内のテキストファイルを結合",
	Long:  "任意のフォルダ内のテキストファイルを結合",
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
		fmt.Printf("%s に結合したデータを保存しました。\n", filePath)
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
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// フラグの値を変数にバインド
	rootCmd.Flags().StringVar(&path, "path", "", "ファイル保存場所")
	rootCmd.Flags().IntVar(&row, "row", 1, "ヘッダー行数")

	// 必須のフラグに指定
	rootCmd.MarkFlagRequired("path")
}
