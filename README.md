Wrap string in Go string format
============================================================

Sometimes, we need to write a long string with `` ` `` and `"` in go source file.
But it's a little complex to code this.

This is a simple tool for generate this. enjoy it.

## Install

```bash
go get github.com/leeychee/gowrap
```

## Usage

```bash
# example
cat origin.txt | gowrap -p main -v varname > var.go
```
