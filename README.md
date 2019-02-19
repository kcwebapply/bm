# bm
<img  align="right" src="image/imemo.png" width="100px">

[![GoDoc](https://godoc.org/github.com/kcwebapply/imemo?status.svg)](https://godoc.org/github.com/kcwebapply/imemo)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/imemo)
[](https://github.com/gin-gonic/gin/releases)
[![Release](https://img.shields.io/github/release/kcwebapply/imemo.svg?style=flat-square)](https://github.com/kcwebapply/iemo/release)

`bm` is tag-based cli tool for bookmarking.

![sample-demo](image/imemo.gif)

## Installation

### On macOS

```
brew tap kcwebapply/bm
brew install bm
```

## Usage

### show memo list
listing all memo. `imemo a` also available.
```
> imemo all 
----------------------------------------------------------------------------------
| 1| playing tennis with Mike on Next  Tuesday                                   |
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
```

### save new memo
please input what you want to write on memo.
`imemo s` also available.
```
> imemo save "meeting at 13:30"
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo saved!
```

## delete memo
please input memo's Id which you want to delete.

```
> imemo d 2
----------------------------------------------------------------------------------
| 2| meeting at 13:30                                                            |
----------------------------------------------------------------------------------
memo deleted!
```

