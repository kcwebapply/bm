package provider

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	//FileName equals ${home}/bm.txt
	FileName = ""
	//ContentPath equals ${home}/bm
	ContentPath = ""
)

func init() {
	// file for bookmark list
	filePath, _ := homedir.Dir()
	FileName = fmt.Sprintf("%s/%s", filePath, "bm.txt")

	// create directory for page-content save
	contentPath := filePath + "/" + "bm-content"
	err := os.MkdirAll(contentPath, 0777)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}
	ContentPath = contentPath

}
