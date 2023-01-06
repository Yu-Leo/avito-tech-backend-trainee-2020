package httpserver

import (
	"net"
	"strconv"
	"time"
)

// Option -.
type Option func(*Server)

// Port -.
func Port(port int) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", strconv.Itoa(port))
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
