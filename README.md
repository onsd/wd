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