package runtime

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type virtualFS struct {
	path     string
	files    map[string][]byte
	createAt time.Time
}

type virtualFile struct {
	name     string
	data     []byte
	offset   int64
	createAt time.Time
	isDir    bool
}

type virtualFileInfo struct {
	name     string
	size     int64
	createAt time.Time
	isDir    bool
}

func MakeFS(files map[string][]byte, path string) http.FileSystem {
	path = strings.TrimSuffix(path, "/")

	return &virtualFS{path, files, time.Now()}
}

func (fs *virtualFS) Open(name string) (http.File, error) {
	if name == fs.path {
		return &virtualFile{name, nil, 0, fs.createAt, true}, nil
	}

	fileName := strings.TrimPrefix(name, fs.path)

	if data, ok := fs.files[fileName]; ok {
		return &virtualFile{name, data, 0, fs.createAt, false}, nil
	}

	return nil, os.ErrNotExist
}

func (virtualFile) Close() error {
	return nil
}

func (f *virtualFile) Read(p []byte) (n int, err error) {
	if f.offset >= int64(len(f.data)) {
		return 0, io.EOF
	}

	n = copy(p, f.data[f.offset:])
	f.offset += int64(n)

	return n, nil
}

func (f *virtualFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case 0:
		f.offset = offset
	case 1:
		f.offset += offset
	case 2:
		f.offset = int64(len(f.data)) + offset
	}

	return f.offset, nil
}

func (virtualFile) Readdir(_ int) ([]os.FileInfo, error) {
	return nil, os.ErrNotExist
}

func (f *virtualFile) Stat() (os.FileInfo, error) {
	return &virtualFileInfo{f.name, int64(len(f.data)), f.createAt, f.isDir}, nil
}

func (fi *virtualFileInfo) Name() string {
	return fi.name
}

func (fi *virtualFileInfo) Size() int64 {
	return fi.size
}

func (virtualFileInfo) Mode() os.FileMode {
	return 0
}

func (fi *virtualFileInfo) ModTime() time.Time {
	return fi.createAt
}

func (fi *virtualFileInfo) IsDir() bool {
	return fi.isDir
}

func (virtualFileInfo) Sys() any {
	return nil
}
