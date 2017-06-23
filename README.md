# brightnesses

Get the brightness of images and rename the images based on their brightnesses. This may be useful for eliminating duplicate images in a directory full of images.

# Installation

First install [Go](https://golang.org/doc/install) and configure a GOPATH. Make sure `$GOPATH/bin` is in your PATH. Now run:

```
$ go get github.com/unixpickle/brightnesses
```

# Usage

To get the brightness of a single image:

```
$ brightnesses foo.png
133 foo.png
```

To list all the images in a directory with their corresponding brightnesses:

```
$ brightnesses /path/to/dir
133 /path/to/dir/foo.png
183 /path/to/dir/bar.jpeg
...
```

This can be piped into `sort` on UNIX:

```
$ brightnesses /path/to/dir | sort -n
5 /path/to/dir/baz.png
133 /path/to/dir/foo.png
183 /path/to/dir/bar.jpeg
...
```

You can rename all of the files to have a name prefixed with the brightness followed by a space:

```
$ brightnesses -rename /path/to/dir
133 /path/to/dir/foo.png
183 /path/to/dir/bar.jpeg
...
$ ls /path/to/dir
133 foo.png      183 bar.jpeg      5 baz.png
...
```
