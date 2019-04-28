package repository

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	fileName    = ""
	maxTextSize = 60

	contentPath = ""
)

func init() {

	// file for bookmark list
	filePath, _ := homedir.Dir()
	fileName = fmt.Sprintf("%s/%s", filePath, "bm.txt")

	// create directory for page-content save
	contentPath := filePath + "/" + "bm-content"
	err := os.MkdirAll(contentPath, 0777)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}
}
