package middlewares

import (
	"../config"
	"fmt"
	"net"
)

func Apply(middleware config.MiddlewareConfig, clientConn, backendConn net.Conn) (net.Conn, net.Conn, error) {
	switch middleware.Kind {
	case "min_http":
		return minHttp(middleware, clientConn, backendConn)
		//	case "via":
		//		return via(middleware, clientConn, backendConn)
	default:
		return nil, nil, fmt.Errorf("Unknown middleware kind %s", middleware.Kind)
	}
}
