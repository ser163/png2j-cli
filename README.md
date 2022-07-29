# png2j-cli

## Description
A simple command-line tool for converting PNG files to JPG files.

## Installation
```shell
$ git clone https://github.com/ser163/png2j-cli.git
```
```shell
$ cd png2j-cli
```

```shell
$ go mod download
```

```shell
$ go build -o png2j
```
## Usage
```shell

Usage:
  png2j [flags]

Flags:
  -e, --height uint     height of output JPG file
  -h, --help            help for png2j
  -o, --output string   output JPG file
  -q, --quality uint8   quality value: 1-100 (default 100)
  -s, --source string   source PNG file
  -w, --width uint      width of output JPG file

```
Converting a PNG image:
```shell
$ png2j -s ./image/go.png -o ./image/go.png.jpg
```
Change JPG image size
```shell
$ png2j -s ./image/go.png -o ./image/go.png.jpg -w 948 -e 418
```

Changing the image quality
```shell
$ png2j -s ./image/go.png -o ./image/go.png.jpg -q 90
```