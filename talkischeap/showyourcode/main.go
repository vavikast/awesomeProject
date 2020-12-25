package showyourcode

import (
	"os"
	"path/filepath"
)

var RootDir  string
func inferDir() {
	var infer  func(d string) string
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	infer = func(d string) string {
		if esists(d+"/template"){
			return d
		}
		return infer(filepath.Dir(d))
	}
	RootDir = infer(dir)
}

func esists(d string) bool  {
	_, err := os.Stat(d)
	return err == nil || os.IsExist(err)

}

