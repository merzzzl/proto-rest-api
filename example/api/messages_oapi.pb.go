// Code generated by protoc-gen-go-oapi. DO NOT EDIT.
// versions:
// - protoc-gen-go-oapi v0.0.0-alpha.0
// - protoc             v3.21.12
// source: example/proto/messages.proto

package api

import (
	runtime "github.com/merzzzl/proto-rest-api/runtime"
	swagger "github.com/merzzzl/proto-rest-api/swagger"
	http "net/http"
	strings "strings"
)

// RegisterSwaggerUIHandler registers swagger ui handler.
func RegisterSwaggerUIHandler(mux runtime.ServeMuxer, path string) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	fs, err := swagger.GetSwaggerUI(swaggerJSON, path)
	if err != nil {
		return err
	}

	mux.Handle(path, http.FileServer(fs))

	return nil
}

// RegisterReDocUIHandler registers redoc ui handler.
func RegisterReDocUIHandler(mux runtime.ServeMuxer, path string) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	fs, err := swagger.GetReDocUI(swaggerJSON, path)
	if err != nil {
		return err
	}

	mux.Handle(path, http.FileServer(fs))

	return nil
}

var swaggerJSON = []byte{
	0x7b, 0x22, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x22, 0x3a, 0x22, 0x33, 0x2e, 0x30, 0x2e,
	0x33, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x3a, 0x22, 0x41, 0x70, 0x69, 0x22, 0x2c, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x35, 0x2d, 0x30, 0x32, 0x2d, 0x30, 0x34, 0x22, 0x7d,
	0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x5d,
	0x2c, 0x22, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x5d,
	0x2c, 0x22, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x7d, 0x22, 0x3a, 0x7b, 0x22, 0x67, 0x65, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3a, 0x22, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x45, 0x63, 0x68, 0x6f, 0x22, 0x2c, 0x22, 0x74, 0x61, 0x67,
	0x73, 0x22, 0x3a, 0x5b, 0x22, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x5d, 0x2c, 0x22, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a,
	0x5b, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x22, 0x3a, 0x22, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x22,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a,
	0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x22, 0x7d, 0x5d, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3a, 0x22, 0x41, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x22, 0x2c,
	0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a,
	0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x7d, 0x2c,
	0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x32, 0x30,
	0x30, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3a, 0x22, 0x41, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x22, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x2c, 0x22,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x67, 0x65, 0x74, 0x22,
	0x3a, 0x7b, 0x22, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22, 0x4c, 0x49, 0x53,
	0x54, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x22, 0x2c, 0x22, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3a, 0x22, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x2c, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3a,
	0x5b, 0x22, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x5d, 0x2c, 0x22, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a,
	0x5b, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x2c,
	0x22, 0x69, 0x6e, 0x22, 0x3a, 0x22, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x2c, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x22, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22,
	0x3a, 0x22, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x7d, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x50, 0x61, 0x67, 0x65, 0x20, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x7d, 0x2c, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x22, 0x3a, 0x22,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a,
	0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x22, 0x7d, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3a, 0x22, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x20, 0x70, 0x65, 0x72, 0x20, 0x70, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x2c, 0x7b, 0x22,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x64, 0x73, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x22,
	0x3a, 0x22, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x22, 0x2c, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x7d, 0x7d, 0x2c, 0x22,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x4c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x49, 0x44,
	0x73, 0x22, 0x7d, 0x5d, 0x2c, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22,
	0x3a, 0x7b, 0x22, 0x32, 0x30, 0x30, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x41, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x22, 0x2c,
	0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a,
	0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x7d, 0x7d, 0x7d,
	0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x20, 0x6e, 0x65, 0x77, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x22, 0x2c, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3a, 0x22, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2c, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5d, 0x2c, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x41, 0x20, 0x4a, 0x53, 0x4f,
	0x4e, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x22, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22,
	0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x7d, 0x2c, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x32, 0x30, 0x30, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x41, 0x20, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x22, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22,
	0x3a, 0x7b, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65,
	0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x22, 0x3a, 0x7b, 0x22,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x22, 0x3a, 0x22, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x22, 0x2c, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x3a, 0x22, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x2c, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x5b, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x22, 0x3a, 0x22, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72,
	0x75, 0x65, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x7d, 0x2c,
	0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x49, 0x44, 0x22, 0x7d, 0x5d, 0x2c, 0x22, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x32, 0x30, 0x30, 0x22, 0x3a,
	0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22,
	0x41, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x20, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x67, 0x65, 0x74, 0x22,
	0x3a, 0x7b, 0x22, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22, 0x47, 0x45, 0x54,
	0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x22, 0x2c, 0x22, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3a, 0x22, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x2c, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5d, 0x2c, 0x22,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x5b, 0x7b, 0x22, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x22, 0x3a, 0x22,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22,
	0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b,
	0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22,
	0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x22, 0x7d, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3a, 0x22, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x49, 0x44, 0x22, 0x7d, 0x5d, 0x2c,
	0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x32, 0x30,
	0x30, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3a, 0x22, 0x41, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x22, 0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a, 0x22,
	0x50, 0x41, 0x54, 0x43, 0x48, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x22, 0x2c, 0x22, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3a, 0x22, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x3a, 0x5b, 0x22, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x3a, 0x5b, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22,
	0x69, 0x6e, 0x22, 0x3a, 0x22, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x7d, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x49,
	0x44, 0x22, 0x7d, 0x5d, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3a, 0x22, 0x41, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x22,
	0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b,
	0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22,
	0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x7d, 0x7d, 0x2c,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x7d,
	0x2c, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x32,
	0x30, 0x30, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3a, 0x22, 0x41, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22,
	0x70, 0x75, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3a,
	0x22, 0x50, 0x55, 0x54, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x22, 0x2c, 0x22, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3a, 0x22, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x50, 0x75, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3a, 0x5b, 0x22,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5d,
	0x2c, 0x22, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x5b, 0x7b,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x22,
	0x3a, 0x22, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22,
	0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x22, 0x7d, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3a, 0x22, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x49, 0x44, 0x22, 0x7d,
	0x5d, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x3a,
	0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22,
	0x41, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x22, 0x2c, 0x22, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23,
	0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x50, 0x75, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x7d, 0x2c, 0x22, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x32, 0x30, 0x30, 0x22, 0x3a, 0x7b,
	0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x41,
	0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x20, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x22, 0x7d, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x7d, 0x22, 0x3a, 0x7b, 0x22, 0x67, 0x65, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3a, 0x22, 0x45, 0x63, 0x68, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x2c, 0x22,
	0x74, 0x61, 0x67, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x22, 0x3a, 0x5b, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x2c, 0x22, 0x69, 0x6e, 0x22, 0x3a, 0x22, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75, 0x65, 0x2c,
	0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x7d, 0x2c, 0x22, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x7d, 0x5d, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3a, 0x22, 0x41, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x22,
	0x2c, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b,
	0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22,
	0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7d, 0x7d,
	0x7d, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x74, 0x72, 0x75,
	0x65, 0x7d, 0x2c, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x3a, 0x7b,
	0x22, 0x32, 0x30, 0x30, 0x22, 0x3a, 0x7b, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x41, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66,
	0x75, 0x6c, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x22, 0x2c, 0x22, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23,
	0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7d, 0x7d, 0x7d, 0x7d,
	0x7d, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x3a, 0x7b, 0x22, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x22, 0x7d, 0x2c, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x7b, 0x22,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d,
	0x7d, 0x2c, 0x22, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b,
	0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x3a, 0x7b, 0x22, 0x6b, 0x65, 0x79, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a,
	0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x7d, 0x2c, 0x22, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x22, 0x2c, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x24,
	0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x7d, 0x2c, 0x22,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22,
	0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x7d,
	0x2c, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x65, 0x6e, 0x75, 0x6d, 0x22, 0x3a, 0x5b, 0x22,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x22, 0x2c, 0x22, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x41, 0x46,
	0x54, 0x22, 0x2c, 0x22, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49,
	0x53, 0x48, 0x45, 0x44, 0x22, 0x5d, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a,
	0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x2c, 0x22, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23,
	0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x2c, 0x22, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22,
	0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x50, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x3a, 0x7b, 0x22, 0x24, 0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x7d, 0x2c, 0x22, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x3a, 0x7b, 0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x7d, 0x7d, 0x7d,
	0x2c, 0x22, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a,
	0x7b, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3a, 0x7b, 0x22, 0x24,
	0x72, 0x65, 0x66, 0x22, 0x3a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x7d, 0x7d,
	0x7d, 0x2c, 0x22, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3a, 0x7b, 0x22,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a,
	0x22, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x7d, 0x2c, 0x22, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x22, 0x7d, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x3a, 0x5b, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x45, 0x63, 0x68, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x7d, 0x2c, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3a, 0x22, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x7d, 0x5d, 0x7d,
}
