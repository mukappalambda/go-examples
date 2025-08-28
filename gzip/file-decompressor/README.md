# Decompressing gzip files

```console
$ go build
$ ./file-decompressor -h
Usage of ./file-decompressor:
  -in string
        input file
```

Create an example `.gz` file called `my-file.txt.gz`:

```console
$ cat << EOF > my-file.txt
gzip is a file format and a software application used for file compression and decompression. The program was created by Jean-loup Gailly and Mark Adler as a free software replacement for the compress program used in early Unix systems, and intended for use by GNU (from which the "g" of gzip is derived). Version 0.1 was first publicly released on 31 October 1992, and version 1.0 followed in February 1993.
EOF
$ gzip my-file.txt
```

Decompress the gzip file and read its content:

```console
$ ./file-decompressor -in my-file.txt.gz
gzip is a file format and a software application used for file compression and decompression. The program was created by Jean-loup Gailly and Mark Adler as a free software replacement for the compress program used in early Unix systems, and intended for use by GNU (from which the "g" of gzip is derived). Version 0.1 was first publicly released on 31 October 1992, and version 1.0 followed in February 1993.
```

To clean up:

```console
$ rm -f my-file.txt.gz
```
