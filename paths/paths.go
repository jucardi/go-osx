package paths

import (
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

var created []string

// Indicates whether a file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Creates a temporary directory and returns the location
func TempDir() (ret string, err error) {
	ret = path.Join(os.TempDir(), uuid.New().String())
	err = os.Mkdir(ret, 0755)
	addDir(ret)
	return
}

func Combine(elem ...string) string {
	p := strings.Replace(path.Join(elem...), ":/", "://", -1)

	if !strings.Contains(p, "\\") {
		return p
	}

	p1 := strings.Replace(p, "/", "\\", -1)
	p2 := strings.Replace(p1, "\\\\", "\\", -1)
	return strings.Replace(p2, "\\\\", "\\", -1)
}

func addDir(dir string) {
	if dir == "" {
		return
	}
	created = append(created, dir)
}
