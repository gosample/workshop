package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/gosample/workshop/proto/users"
	"github.com/gosample/workshop/shared"
	"github.com/gosample/workshop/users/server"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	srv := grpc.NewServer()
	users.RegisterUsersServer(srv, &server.Server{})
	srv.Serve(listener)
}
