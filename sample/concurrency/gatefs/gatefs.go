package gatefs

import (
	"golang.org/x/tools/godoc/vfs"
	"os"
)

type gatefs struct {
	fs vfs.FileSystem
	gate
}

func (fs gatefs) Lstat(p string) (os.FileInfo, error) {
	fs.enter()
	defer fs.leave()
	return fs.fs.Lstat(p)
}

type gate chan bool

func (g gate) enter() { g <- true }
func (g gate) leave() { <-g }
