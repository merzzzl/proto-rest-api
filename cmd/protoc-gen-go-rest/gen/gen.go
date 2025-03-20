package gen

import "google.golang.org/protobuf/compiler/protogen"

const (
	contextPackage = protogen.GoImportPath("context")
	ioPackage      = protogen.GoImportPath("io")
	httpPackage    = protogen.GoImportPath("net/http")
	jsonPackage    = protogen.GoImportPath("encoding/json")
	stringsPackage = protogen.GoImportPath("strings")
	runtimePackage = protogen.GoImportPath("github.com/merzzzl/proto-rest-api/runtime")
)
