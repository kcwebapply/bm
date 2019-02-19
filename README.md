# bm
<img  align="right" src="image/imemo.png" width="100px">

[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/imemo)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/imemo)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/imemo.svg?style=flat-square)](https://github.com/kcwebapply/iemo/release)

`bm` is tag-based cli tool for bookmarking.

![sample-demo](bm/bm.gif)

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



By `bm ls ${searchWord}` , you can filtering bookmark by title.


By `bm -t ${tagName}` , you can search bookmark by tag name.



### add bookmark
`bm add ${url}`  `bm a ${url}` 
you can save bookmark by `bm add` command.

you can set tags (max 3 tags) on your bookmark.


## delete bookmark
`bm delete ${bookmark-id}` `bm d ${bookmark-id}`


To delete bookmark on you list, please input bookmark-id that automatically assigned to all bookmark.

you can confirm bookmark-id by `bm ls` command.

```

```



