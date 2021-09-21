package dns

import "net"

// WriteToSessionUDP acts just like net.UDPConn.WriteTo(), but uses a *SessionUDP instead of a net.Addr.
func WriteToSessionUDP(conn *net.UDPConn, b []byte, session *SessionUDP) (int, error) {
	oob := correctSource(session.context)
	n, _, err := conn.WriteMsgUDP(b, oob, session.raddr)
	return n, err
}
