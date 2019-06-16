package repository

import (
	"fmt"

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
	_, _ = sess.Exec("create table page(id INTEGER PRIMARY KEY AUTOINCREMENT, url TEXT, title TEXT,tags TEXT,content TEXT);")
}

//GetPages returns all page entities user saved.
func GetPages() ([]model.Page, error) {
	var rows []model.Page
	_, err := sess.Select("*").From("page").Load(&rows)
	return rows, err
}

// GetPagesByTitleWordGrep retunrs page grepped title by input word.
func GetPagesByTitleWordGrep(word string) ([]model.Page, error) {
	var rows []model.Page
	likeItem := fmt.Sprintf("%%%s%%", word)
	_, err := sess.Select("*").From("page").Where("title like ?", likeItem).Load(&rows)
	return rows, err
}

//GetPagesByTag returns page entities grepped by input tag word.
func GetPagesByTag(tag string) ([]model.Page, error) {
	var rows []model.Page
	likeItem := fmt.Sprintf("%%%s%%", tag)
	_, err := sess.Select("*").From("page").Where("tags like ?", likeItem).Load(&rows)
	return rows, err
}

// GetPagesByContentSearch returns page entities grepped by html-content-search.
func GetPagesByContentSearch(word string) ([]model.Page, error) {
	var rows []model.Page
	likeItem := fmt.Sprintf("%%%s%%", word)
	_, err := sess.Select("*").From("page").Where("content like ?", likeItem).Load(&rows)
	return rows, err
}

// AddPage saved bookrmark user input.
func AddPage(newPage model.Page) error {
	_, err := sess.InsertInto("page").Columns("url", "title", "tags", "content").Record(newPage).Exec()
	return err
}

// RemovePage remove bookmark
func RemovePage(id string) error {
	if _, err := sess.DeleteFrom("page").Where("id = ?", id).Exec(); err != nil {
		return err
	}
	return nil
}
