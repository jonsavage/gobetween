package middlewares

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"strings"
	"sync"
)

func newClientHeaderAdder(clientConn net.Conn, headers map[string]string) net.Conn {
	return &clientHeaderAdder{
		Conn:    clientConn,
		headers: headers,
		mu:      &sync.Mutex{},
	}
}

type clientHeaderAdder struct {
	net.Conn
	headers map[string]string
	remain  io.Reader
	done    bool
	mu      *sync.Mutex
}

func (x *clientHeaderAdder) Read(buf []byte) (int, error) {
	x.mu.Lock()
	defer x.mu.Unlock()

	if x.done {
		return x.remain.Read(buf)
	}

	backup := &bytes.Buffer{}
	reader := bufio.NewReader(x.Conn)

	done := func(err error) (int, error) {
		x.remain = io.MultiReader(backup, reader)
		x.done = true
		i, err2 := x.remain.Read(buf)
		if err != nil {
			return i, err
		}
		return i, err2
	}

	//skip initial crlf if any
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return done(err)
		}

		if b == '\r' || b == '\n' {
			backup.WriteByte(b)
			continue
		}
		reader.UnreadByte()
		break
	}

	//reading first line
	line, err := reader.ReadString('\n')
	backup.WriteString(line)

	if err != nil {
		return done(err)
	}

	//it looks like not http, so we have nothing to do here
	if !strings.Contains(line, "HTTP") {
		return done(nil)
	}

	//reading all headers
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			backup.WriteString(line)
			return done(err)
		}

		if line == "\r\n" {
			reader.UnreadByte()
			reader.UnreadByte()
			break
		}

		backup.WriteString(line)

	}

	for k, v := range x.headers {
		if k == "" {
			continue
		}
		backup.WriteString(k + ": " + v + "\r\n")
	}

	return done(nil)
}
