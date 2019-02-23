
<img  src="image/bm.png" width="100px">

[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/imemo)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/imemo)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/imemo.svg?style=flat-square)](https://github.com/kcwebapply/iemo/release)

`bm` is tag-based cli tool for bookmarking.

<img src="image/bm-sample.png" width="1000px"/>



## Installation

### On macOS

```
brew tap kcwebapply/bm
brew install bm
```

## Usage

### view bookmarks
`bm list` `bm ls`

listing all book. `bm ls` also available.

<img src="image/ls.gif">



By `bm ls ${searchWord}` , you can filtering bookmark by title.
By `bm -t ${tagName}` , you can search bookmark by tag name.

![sample-demo](image/tag.gif)



### add bookmark
`bm add ${url}`  `bm a ${url}`
you can save bookmark by `bm add` command.

![sample-demo](image/add.gif)

you can set tags (max 3 tags) on your bookmark.


## delete bookmark
`bm delete ${bookmark-id}` `bm d ${bookmark-id}`

![sample-demo](image/del.gif)


To delete bookmark on you list, please input bookmark-id that automatically assigned to all bookmark.

you can confirm bookmark-id by `bm ls` command.

## open bookmark
`bm open ${bookmark-id}`
you can open bookmark from cli.

![sample-demo](image/open.gif)
