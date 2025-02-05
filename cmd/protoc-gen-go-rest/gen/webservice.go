package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func WebService(g *protogen.GeneratedFile, service *protogen.Service, requireUnimplemented bool) {
	WebServiceInterface(g, service, requireUnimplemented)
	UnimplementedStruct(g, service, requireUnimplemented)
}
