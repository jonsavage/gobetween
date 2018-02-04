package middlewares

import (
	"../config"
	"net"
)

/**
 * Should parse first HTTP request headers from client add header X-Forwarded-For with client ip and forward
 * this modified headers to backendConn
 */
func minHttp(conf *config.MiddlewareConfig, clientConn, backendConn net.Conn) (net.Conn, net.Conn, error) {

	xff := getIp(clientConn.RemoteAddr())
	via := getIp(clientConn.LocalAddr())

	return newClientHeaderAdder(clientConn, map[string]string{
		conf.MinimalHttpMiddlewareConfig.XffHeaderName: xff,
		conf.MinimalHttpMiddlewareConfig.ViaHeaderName: via,
	}), backendConn, nil

}

func getIp(addr net.Addr) string {

	tcpAddr, ok := addr.(*net.TCPAddr)
	if !ok {
		return addr.String()
	}

	return tcpAddr.IP.String()

}
