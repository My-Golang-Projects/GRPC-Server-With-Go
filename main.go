package main

import (
	"context"
	"log"
	"net"

	invoicer "github.com/My-Golang-Projects/GRPC-Server-With-Go/invoicer/bufs/v1"
	"google.golang.org/grpc"
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
	// defer ln.Close()

	s.ln = ln

	return nil
}

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServiceServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	server := NewServer(":8080")
	err := server.Start()
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(server.ln)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
