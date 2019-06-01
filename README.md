# GoText
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


$ gotext -f sample.txt
Number of words      :  6
Number of characters :  28
```