package dns

import (
	"net"
	"syscall"
)

// WriteToSessionUDP acts just like net.UDPConn.WriteTo(), but uses a *SessionUDP instead of a net.Addr.
func WriteToSessionUDP(conn *net.UDPConn, b []byte, session *SessionUDP) (int, error) {
	oob := correctSource(session.context)
	// Allow freebind for IPv6 features.
	f, err := conn.File()
	if err != nil {
		return 0, err
	}
	if err := syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IP, syscall.IP_FREEBIND, 1); err != nil {
		return 0, err
	}
	n, _, err := conn.WriteMsgUDP(b, oob, session.raddr)
	return n, err
}
