# go-osx

Some OS utilities for go.

## Getting started

```bash
go get github.com/jucardi/go-osx
```

## Utilities provided in this library

- `paths.Exists`: Indicates whether a file or directory exists.
- `paths.TempDir`: Creates a temporary directory inside the global TempDir for the running OS, using a random UUID as its name, and returns the location
- `paths.Combine`: Combines multiple URL, URI or FilePath using `path.Join`. Handles double `/` or double `\` in a smart way that `path.Join` would not. Eg:

```Go
path.Join("http://", "www.abc.com/", "/something/", "/a")
// returns http://www.abc.com//something//a

paths.Combine("http://", "www.abc.com/", "/something/", "/a")
// returns http://www.abc.com/something/a   (Corrects the sets of "//")
```

- `osx.CopyFile`: Copies the contents of the source file to the given destination file. The file will be created if it does not already exist. If the destination file exists, all it's contents will be replaced by the contents of the source file.
