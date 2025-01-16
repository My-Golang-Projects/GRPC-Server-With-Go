package main

import (
	"net"

	"github.com/prometheus/common/log"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
}
