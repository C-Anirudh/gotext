![gitinfo-cover](logo.png)

[![GitHub Watches](https://img.shields.io/github/watchers/C-Anirudh/gotext.svg?style=social&label=Watch&maxAge=2592000)](https://GitHub.com/C-Anirudh/gotext/watchers)
[![GitHub Starts](https://img.shields.io/github/stars/C-Anirudh/gotext.svg?style=social&label=Star&maxAge=2592000)](https://GitHub.com/C-Anirudh/gotext/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/C-Anirudh/gotext.svg?style=social&label=Fork&maxAge=2592000)](https://GitHub.com/C-Anirudh/gotext/network)

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
Usage: gotext [options]
Options:
         -f      The name/path of text file to be parsed
```

```
$ gotext -f sample.txt
Number of words      :  6
Number of characters :  28
```