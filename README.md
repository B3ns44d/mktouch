# mktouch

`mktouch` is a command-line tool for creating a file at a specified path and, if necessary, creating any parent directories that don't already exist.

## Installation

To install `mktouch`, run the following command:

```bash
go get -u github.com/B3ns44d/mktouch
```


## Usage

To create an empty file at a specified path, run the following command:

```bash
mktouch <filename>
```

Replace `<filename>` with the path to the file you want to create.

By default, `mktouch` will create the file with read and write permissions for the owner, and read-only permissions for others. You can use the `--mode` flag to specify different permissions in octal format:

```bash
mktouch --mode 0640 <filename>
```

This will create the file with read and write permissions for the owner, and read-only permissions for the group.

You can also use the `--parents` or `-p` flag to create parent directories if they don't already exist:

```
mktouch -p /path/to/new/file
```

This will create the directories `/path/to/new/` if they don't already exist, and then create the file `file` in the `new/` directory.

You can create multiple files at once by passing in multiple filenames as arguments:


```bash
mktouch -p /path/to/new/{file.go,file1.go,file2.go}
```

This will create the files file.go, file1.go, and file2.go in the `/path/to/new/` directory.

You can also use brace expansion to create a range of files:

```bash
mktouch -p /path/to/new/test{0001..0003}.go
```

This will create the files test0001.js, test0002.js, and test0003.js in the `/path/to/new/` directory. Note that brace expansion is a feature of the shell, not the mktouch command itself, so you need to run these commands in a shell that supports brace expansion.


