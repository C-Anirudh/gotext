![gitinfo-cover](logo.png)

[![GitHub Watches](https://img.shields.io/github/watchers/C-Anirudh/gotext.svg?style=social&label=Watch&maxAge=2592000)](https://github.com/C-Anirudh/gotext/watchers)
[![GitHub Starts](https://img.shields.io/github/stars/C-Anirudh/gotext.svg?style=social&label=Star&maxAge=2592000)](https://github.com/C-Anirudh/gotext/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/C-Anirudh/gotext.svg?style=social&label=Fork&maxAge=2592000)](https://github.com/C-Anirudh/gotext/network)

CLI tool to parse text files and get count of words and characters.


## :minidisc: Installation
```
$ go get github.com/C-Anirudh/gotext
```

To be able to run the command from anywhere in the computer, run the following commands
```
$ cd $GOPATH/bin/
$ sudo cp gotext /usr/bin/gotext
```

## :rocket: Usage
```
$ gotext --help
```
OR
```
$ gotext
Usage: ./gotext [options]
Options:
         -w      absolute path to text file for count of words
         -c      absolute path to text file for count of characters
```

### Examples
```
$ gotext -w sample.txt
6
```
```
$ gotext -c sample.txt
28
```
```
$ gotext -w sample.txt -c sample.txt
6
28
```