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
	files, err := swaggerUI.ReadDir("assets/swagger")
	if err != nil {
		return nil, fmt.Errorf("read assets/swagger directory: %w", err)
	}

	fileData := make(map[string][]byte)

	for _, file := range files {
		filePath := filepath.Join("assets/swagger", file.Name())

		data, err := swaggerUI.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("read file: %w", err)
		}

		fileData["/"+file.Name()] = data
	}

	fileData["/swagger.json"] = spec

	return runtime.MakeFS(fileData, path), nil
}

func GetReDocUI(spec []byte, path string) (http.FileSystem, error) {
	files, err := swaggerUI.ReadDir("assets/redoc")
	if err != nil {
		return nil, fmt.Errorf("read assets/redoc directory: %w", err)
	}

	fileData := make(map[string][]byte)

	for _, file := range files {
		filePath := filepath.Join("assets/redoc", file.Name())

		data, err := swaggerUI.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("read file: %w", err)
		}

		fileData["/"+file.Name()] = data
	}

	fileData["/swagger.json"] = spec

	return runtime.MakeFS(fileData, path), nil
}
