package main

import (
	"net"

	invoicer "github.com/My-Golang-Projects/GRPC-Server-With-Go/invoicer/bufs/v1"
	"github.com/prometheus/common/log"
)

type Server struct {
	listenAddr string
	ln         net.Listener
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	s.ln = ln

	return nil
}

func main() {
	err := NewServer(":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	// invoicer.New
}
