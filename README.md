# dup

dup finds duplicate files in a directory hierarchy (recursive) based
on their SHA256 hash.

Digests are cached in `hashes.dat`. You can add/delete files and run
dup again. However, it is assumed that file contents don't change
between dup runs. If you want dup to re-build its index, remove
`hashes.dat`. You can also remove individual lines from `hashes.dat`
to make dup re-calculate specific digests.
