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

### use homebrew
```console
$ brew install onsd/tools/wd
```

### download binary from github releases

check [releases](https://github.com/onsd/wd/releases) and download assets for your environment

### clone and install locally

**go required**
 ```console
$ git clone git@github.com:onsd/wd.git
$ cd wd
$ go install
```
