package swagger

import (
	"embed"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/merzzzl/proto-rest-api/runtime"
)

//go:embed assets/*
var swaggerUI embed.FS

func GetSwaggerUI(spec []byte, path string) (http.FileSystem, error) {
	files, err := swaggerUI.ReadDir("assets")
	if err != nil {
		return nil, fmt.Errorf("read assets directory: %w", err)
	}

	fileData := make(map[string][]byte)

	for _, file := range files {
		filePath := filepath.Join("assets", file.Name())

		data, err := swaggerUI.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("read file: %w", err)
		}

		fileData["/"+file.Name()] = data
	}

	fileData["/swagger.json"] = spec

	return runtime.MakeFS(fileData, path), nil
}
