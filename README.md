# snappy-cli
Simple CLI utility for the snappy compression format

Contains two complementary CLI utilities, one for compressing, the other
for uncompressing i.e `snappy-compress` and `snappy-uncompress`.

* Note:
Both utils:
- Read either from stdin or from a file on disk passed in as a path.
- Output their result to stdout.


### Installation and usage

- snappy-compress

CLI utility to compress data in the snappy format.
It takes in either a source file or content from stdin e.g

```shell
$ go get github.com/odeke-em/snappy-cli/cmd/snappy-compress
$ cat content | snappy-compress > output.compressed
$ snappy-compress ~/Downloads/fileUp > output.compressed
```

- snappy-uncompress

```shell
$ go get github.com/odeke-em/snappy-cli/cmd/snappy-uncompress
$ cat content.compressed | snappy-uncompress > uncompressed
$ snappy-uncompress content.compressed > uncompressed
```
