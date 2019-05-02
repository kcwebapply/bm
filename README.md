
<img  src="https://imgur.com/UM1Reoh.png" width="100px" align="right">

# bm


[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/bm)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/bm)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/bm.svg?style=flat-square)](https://github.com/kcwebapply/bm/release)

you can bookmark webpage and search by `title`, `tag`, `html-body`.


## Installation

### On macOS

```
brew tap kcwebapply/bm
brew install bm
```

## Usage

### view bookmarks

#### listing bookmarks
`bm list` `bm ls`
<img src="https://imgur.com/jJdjTAU.png">

###### content-search
By `bm ls -s ${word}`, you can search in html contents search.
<img src="https://imgur.com/e2TdtjZ.png">

###### title-grep
`bm ls ${searchWord}` , you can filtering bookmark by title.

###### tag-search
`bm ls -t ${tagName}` , you can search bookmark by tag name.

<img src="https://imgur.com/wsVVaOA.png">



#### add bookmark
`bm add ${url}`  `bm a ${url}`
you can save bookmark by `bm add` command.

<img src="https://imgur.com/fT3dRDk.png">

you can set tags (max 3 tags) on your bookmark.

###### import chrome bookmark
`bm import {filepath}`
you can import chrome bookmark export file.

<img src="https://imgur.com/CwbWbQc.gif">


#### delete bookmark
`bm rm ${bookmark-id}`

To delete bookmark on you list, please input bookmark-id that automatically assigned to all bookmark.

you can confirm bookmark-id by `bm ls` command.

#### open bookmark
`bm open ${bookmark-id}`
you can open bookmark web-site from cli.
