package openapi

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func FormatedPath(service *protogen.Service, method *protogen.Method) (string, error) {
	serviceOptions, ok := service.Desc.Options().(*descriptorpb.ServiceOptions)
	if !ok {
		return "", fmt.Errorf("unknown service options in %s", service.GoName)
	}

	extValSrv := proto.GetExtension(serviceOptions, restapi.E_Service)

	serviceRule, ok := extValSrv.(*restapi.ServiceRule)
	if !ok {
		return "", fmt.Errorf("unknown http options in %s", service.GoName)
	}

	methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return "", fmt.Errorf("unknown method options in %s", method.GoName)
	}

	extVal := proto.GetExtension(methodOptions, restapi.E_Method)

	restRule, ok := extVal.(*restapi.MethodRule)
	if !ok {
		return "", fmt.Errorf("unknown http options in %s", method.GoName)
	}

	basePath := serviceRule.GetBasePath()
	basePath = strings.TrimSuffix(basePath, "/")

	path := strings.Split(restRule.GetPath(), "?")[0]
	segs := strings.Split(path, "/")

	for i, seg := range segs {
		if strings.HasPrefix(seg, ":") {
			segs[i] = "{" + seg[1:] + "}"
		}
	}

	return basePath + "/" + strings.Join(segs, "/"), nil
}
