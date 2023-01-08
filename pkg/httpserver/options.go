package httpserver

import (
	"net"
	"strconv"
	"time"
)

type Option func(*Server)

func HostPort(host string, port int) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort(host, strconv.Itoa(port))
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
