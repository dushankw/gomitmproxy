package gomitmproxy

import "net"

// ClientConn returns the downstream client <-> proxy connection
func (s *Session) ClientConn() net.Conn {
    return s.conn
}

// ServerConn returns the upstream connection
func (s *Session) ServerConn() net.Conn {
    return s.serverConn
}
