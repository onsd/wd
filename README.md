# wd
wd - a command line tool for print specified position of words

## usage
```console
$ echo "a b c" | wd -n 1
"a"
```
```console
$ cat some-text
a b c
d e f

$ wd -n 1 some-text
a
d
```

## how to install
1. clone and install locally
**go required**
 ```console
$ git clone git@github.com:onsd/wd.git
$ cd wd
$ go install
```

2. download binary from github releases

check [releases](https://github.com/onsd/wd/releases) and download assets for your environment