package gen

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func WebSocket(g *protogen.GeneratedFile, service *protogen.Service) error {
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			return nil
		}

		WebSocketInreface(g, service, method)

		g.P()

		if err := WebSocketStruct(g, service, method); err != nil {
			return fmt.Errorf("failed to generate websocket struct for %s: %w", method.GoName, err)
		}

		g.P()

		if err := WebSocketFuncs(g, service, method); err != nil {
			return fmt.Errorf("failed to generate websocket funcs for %s: %w", method.GoName, err)
		}

		g.P()
	}

	return nil
}
