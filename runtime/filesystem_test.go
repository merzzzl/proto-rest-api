package runtime_test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestFS_0(t *testing.T) {
	t.Parallel()

	files := map[string][]byte{
		"/index.html": []byte("Hello, World!"),
		"/style.css":  []byte("body { color: red; }"),
	}

	fs := runtime.MakeFS(files, "/static")

	file, err := fs.Open("/static/index.html")
	require.NoError(t, err)

	info, err := file.Stat()
	require.NoError(t, err)

	require.Equal(t, "/static/index.html", info.Name())
	require.Equal(t, int64(13), info.Size())
	require.False(t, info.IsDir())
	require.Zero(t, info.Mode())
	require.Nil(t, info.Sys())
	require.NotEqual(t, time.Now(), info.ModTime())

	data := make([]byte, 13)

	n, err := file.Read(data)
	require.NoError(t, err)

	require.Equal(t, 13, n)
	require.Equal(t, []byte("Hello, World!"), data)

	err = file.Close()
	require.NoError(t, err)
}

func TestFS_1(t *testing.T) {
	t.Parallel()

	files := map[string][]byte{
		"/index.html": []byte("Hello, World!"),
	}

	fs := runtime.MakeFS(files, "/static")

	_, err := fs.Open("/static/style.css")
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestFS_2(t *testing.T) {
	t.Parallel()

	files := map[string][]byte{
		"/index.html": []byte("Hello, World!"),
	}

	fs := runtime.MakeFS(files, "/static")

	file, err := fs.Open("/static")
	require.NoError(t, err)

	info, err := file.Stat()
	require.NoError(t, err)

	require.Equal(t, "/static", info.Name())
	require.Equal(t, int64(0), info.Size())
	require.True(t, info.IsDir())

	_, err = file.Readdir(0)
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestFS_3(t *testing.T) {
	t.Parallel()

	files := map[string][]byte{
		"/index.html": []byte("Hello, World!"),
	}

	fs := runtime.MakeFS(files, "/static")

	file, err := fs.Open("/static/index.html")
	require.NoError(t, err)

	buf := make([]byte, 5)

	n, err := file.Read(buf)
	require.NoError(t, err)

	require.Equal(t, 5, n)
	require.Equal(t, []byte("Hello"), buf)

	_, err = file.Seek(0, 0)
	require.NoError(t, err)

	n, err = file.Read(buf)
	require.NoError(t, err)

	require.Equal(t, 5, n)
	require.Equal(t, []byte("Hello"), buf)

	_, err = file.Seek(2, 1)
	require.NoError(t, err)

	n, err = file.Read(buf)
	require.NoError(t, err)

	require.Equal(t, 5, n)
	require.Equal(t, []byte("World"), buf)

	_, err = file.Seek(1, 2)
	require.NoError(t, err)

	_, err = file.Read(buf)
	require.ErrorIs(t, err, io.EOF)
}
