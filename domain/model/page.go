package model

// Page type is struct of one webpage data.
// Id is primary key of bookmark page.
// Title is webpage's title  and Content is webpage's html body. Both of them got by http request.
type Page struct {
	ID      int
	URL     string
	Title   string
	Tags    string
	Content string
}
