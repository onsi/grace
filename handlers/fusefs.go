package handlers

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/hanwen/go-fuse/splice"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

const fusemount = "/tmp/fusefs-mount"

type MountFUSEFS struct {
}

func (p *MountFUSEFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(fusemount)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("already mounted"))
		return
	}
	err = os.Mkdir(fusemount, 0755)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("mkdir fail: " + err.Error()))
	}

	nfs := pathfs.NewPathNodeFs(&HelloFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(fusemount, nfs.Root(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("mount fail: " + err.Error()))
	}
	go server.Serve()
}

type ListFUSEFS struct {
}

func (p *ListFUSEFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(fusemount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("list fail: " + err.Error()))
	}
	for _, file := range files {
		w.Write([]byte(file.Name() + "\n"))
	}
}

//from https://github.com/hanwen/go-fuse/blob/master/example/hello/main.go
type HelloFs struct {
	pathfs.FileSystem
}

func (me *HelloFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	switch name {
	case "fuse-fs-works.txt":
		return &fuse.Attr{
			Mode: fuse.S_IFREG | 0644, Size: uint64(len(name)),
		}, fuse.OK
	case "":
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (me *HelloFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	if name == "" {
		c = []fuse.DirEntry{{Name: "fuse-fs-works.txt", Mode: fuse.S_IFREG}}
		return c, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (me *HelloFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	if name != "fuse-fs-works.txt" {
		return nil, fuse.ENOENT
	}
	if flags&fuse.O_ANYWRITE != 0 {
		return nil, fuse.EPERM
	}
	return nodefs.NewDataFile([]byte(name)), fuse.OK
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello MOUNTPOINT")
	}
	nfs := pathfs.NewPathNodeFs(&HelloFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(flag.Arg(0), nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Serve()
}
