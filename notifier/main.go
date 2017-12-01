package main

import (
	"log"
	"net"

	"github.com/gosample/workshop/notifier/server"
	"github.com/gosample/workshop/proto/notifier"
	"github.com/gosample/workshop/shared"
	"google.golang.org/grpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	srv := grpc.NewServer()
	notifier.RegisterNotifierServer(srv, &server.Server{})
	srv.Serve(listener)
}
