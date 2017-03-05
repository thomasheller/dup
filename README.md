# dup

[![Go Report Card](https://goreportcard.com/badge/github.com/thomasheller/dup)](https://goreportcard.com/report/github.com/thomasheller/dup)

dup finds duplicate files in a directory hierarchy (recursive) based
on their SHA256 hash.

Digests are cached in `hashes.dat`. You can add/delete files and run
dup again. However, it is assumed that file contents don't change
between dup runs. If you want dup to re-build its index, remove
`hashes.dat`. You can also remove individual lines from `hashes.dat`
to make dup re-calculate specific digests.

## Example

```
$ echo foo >{1,2,3}.txt
$ echo bar >4.txt
$ dup
b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c:1.txt
b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c:2.txt
b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c:3.txt
$ cat hashes.dat
1.txt:b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c
2.txt:b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c
3.txt:b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c
4.txt:7d865e959b2466918c9863afca942d0fb89d7c9ac0c99bafc3749504ded97730
```

