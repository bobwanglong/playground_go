package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 穿件文件夹
func Mkdir(path string) {
	os.MkdirAll(path, 0o700)
}

// ExpandDir清除目录，如果是绝对的，则返回自身， //否则以当前工作目录为前缀
func ExpandDir(dir string) string {
	wd := filepath.Dir(os.Args[0])
	if filepath.IsAbs(dir) {
		return filepath.Clean(dir)
	}
	return filepath.Clean(filepath.Join(wd, dir))
}

func main() {
	// Mkdir("mydir")
	path := ExpandDir("/mydir") // "/mydir"
	// path := ExpandDir("mydir") // "/var/folders/zn/nvphnz_j6sd1wnpgj79y3wnm0000gn/T/go-build686380358/b001/exe/mydir"
	fmt.Println(path)
}
