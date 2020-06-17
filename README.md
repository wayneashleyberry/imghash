> command `imghash` is a simple wrapper around the [corona10/goimagehash](https://github.com/corona10/goimagehash) library.

[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/imghash)](https://goreportcard.com/report/github.com/wayneashleyberry/imghash)
![Go](https://github.com/wayneashleyberry/imghash/workflows/Go/badge.svg)

### Installation

```sh
cd $HOME
go get github.com/wayneashleyberry/imghash
```

### Usage

```sh
Usage:
  imghash [file] [flags]

Flags:
  -h, --help   help for imghash
```

### Example

```sh
$ imghash testdata/octocat-de-los-muertos.jpg
WYzVi9sUE2Wn8N7KP5uwjg==
a:c30081c3c3c7e7ff
d:0f2b1717170f0f1f
p:bc6bc2b689a5b0a5
```
