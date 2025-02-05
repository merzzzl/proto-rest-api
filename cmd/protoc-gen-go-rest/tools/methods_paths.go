package tools

import (
	"fmt"
	"strings"

	"github.com/merzzzl/proto-rest-api/restapi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func MethodsPaths(service *protogen.Service) (map[*protogen.Method]*restapi.MethodRule, error) {
	type methodNode struct {
		method map[string]*protogen.Method
		next   map[string]*methodNode
	}

	usedPaths := &methodNode{next: make(map[string]*methodNode), method: make(map[string]*protogen.Method)}
	methods := make(map[*protogen.Method]*restapi.MethodRule, len(service.Methods))

	for _, method := range service.Methods {
		methodOptions, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
		if !ok {
			continue
		}

		extVal := proto.GetExtension(methodOptions, restapi.E_Method)

		restRule, ok := extVal.(*restapi.MethodRule)
		if !ok {
			continue
		}

		path := restRule.GetPath()
		queryParam := make(map[string]struct{})

		sep := strings.LastIndex(path, "?")
		if sep != -1 {
			for _, param := range strings.Split(path[sep+1:], "&") {
				queryParam[param] = struct{}{}
			}

			path = path[:sep]
		}

		path = strings.TrimPrefix(path, "/")
		path = strings.TrimSuffix(path, "/")

		segments := strings.Split(path, "/")
		for i := 1; i < len(segments); i++ {
			if segments[i][0] == ':' {
				if _, ok := queryParam[segments[i]]; ok {
					return nil, fmt.Errorf("%w: mixing query and path parameters (%s) in %s", ErrInvalidPath, segments[i], path)
				}

				segments[i] = "<dynamic>"
			}
		}

		node := usedPaths

		for _, segment := range segments {
			if _, ok := node.next["<dynamic>"]; (ok && segment != "<dynamic>") || (!ok && len(node.next) > 0 && segment == "<dynamic>") {
				return nil, fmt.Errorf("%w: mixing static and dynamic segments in %s", ErrPathAlreadyExists, path)
			}

			if _, ok := node.next[segment]; !ok {
				node.next[segment] = &methodNode{next: make(map[string]*methodNode), method: make(map[string]*protogen.Method)}
			}

			node = node.next[segment]
		}

		if _, ok := node.method[restRule.GetMethod()]; ok {
			return nil, ErrPathAlreadyExists
		}

		node.method[restRule.GetMethod()] = method
		methods[method] = restRule
	}

	return methods, nil
}
