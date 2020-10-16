> command `imghash` is a simple wrapper around the
> [corona10/goimagehash](https://github.com/corona10/goimagehash) and [buckket/go-blurhash](https://github.com/buckket/go-blurhash) libraries.

[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/imghash)](https://goreportcard.com/report/github.com/wayneashleyberry/imghash)
![Go](https://github.com/wayneashleyberry/imghash/workflows/Go/badge.svg)

### Installation

You may [download the latest release](https://github.com/wayneashleyberry/imghash/releases/latest), or compile the binary yourself.

```sh
cd $HOME
go get github.com/wayneashleyberry/imghash
```

### Usage

```sh
Usage:
  imghash [file] [file] [flags]

Flags:
  -h, --help               help for imghash
  -x, --x-components int   blur hash: x components (default 4)
  -y, --y-components int   blur hash: y components (default 3)
```

### Example

Calculate the hash of a single image:

```sh
$ imghash testdata/octocat-de-los-muertos.jpg
WYzVi9sUE2Wn8N7KP5uwjg==
a:c30081c3c3c7e7ff
d:0f2b1717170f0f1f
p:bc6bc2b689a5b0a5
LoPG23nO_NozDis;kCWAwIoLR+W;
```

Calculate the distance between hashes of two images:

```sh
$ imghash testdata/octocat-de-los-muertos.jpg testdata/octocat-de-los-muertos.jpg
a:0
d:0
p:0
```

```sh
$ imghash testdata/octocat-de-los-muertos.jpg testdata/puddle_jumper_octodex.jpg
a:15
d:13
p:28
```
