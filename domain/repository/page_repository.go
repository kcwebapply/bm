package repository

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gocraft/dbr"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/infrastructure/db"
)

var conn *dbr.Connection
var sess *dbr.Session

func init() {
	conn = db.GetConnection()
	sess = conn.NewSession(nil)
	CreatePageTable()
}

//CreatePageTable create page table on sqlite3 db
func CreatePageTable() {
	_, _ = sess.Exec("create table page(id INTEGER PRIMARY KEY, url TEXT, title TEXT,tags TEXT,content TEXT);")
}

//GetPages returns all page entities user saved.
func GetPages() ([]model.Page, error) {
	var rows []model.Page
	_, err := sess.Select("*").From("page").Load(&rows)
	if err != nil {
		fmt.Println("err:", err)
	}
	return rows, err
}

// GetPagesByTitleWordGrep retunrs page grepped title by input word.
func GetPagesByTitleWordGrep(word string) ([]model.Page, error) {
	var rows []model.Page
	_, err := sess.Select("*").From("page").Where("title like ?", word).Load(&rows)
	if err != nil {
		fmt.Println("err:", err)
	}
	return rows, err
}

// AddPage saved bookrmark user input.
func AddPage(newPage model.Page) error {
	_, err := sess.InsertInto("page").Columns("id", "url", "title", "tags", "content").Record(newPage).Exec()
	if err != nil {
		fmt.Println("err:", err)
	}
	return err
}

// RemovePage remove bookmark
func RemovePage(id string) error {
	if _, err := sess.DeleteFrom("page").Where("id = ?", id).Exec(); err != nil {
		fmt.Println("err:", err)
		return err
	}
	return nil
}

func getFileCleanWriter() *bufio.Writer {
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)
	return writer
}