package osx

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func CopyFile(src, dst string, overwrite bool, perm ...os.FileMode) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(dst); !os.IsNotExist(err) && !overwrite {
		return fmt.Errorf("destination file '%s' already exists, overwrite flag set to false, %s", src, err.Error())
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	p := os.FileMode(0655)
	if len(perm) > 0 {
		p = perm[0]
	}
	out, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, p)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
