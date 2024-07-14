package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	grpc_chat "github.com/ArtEmerged/o_chat-server/internal/grpc/chat"
)

func main() {
	l, err := net.Listen("tcp", ":5052")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	grpc_chat.Register(s)

	if err = s.Serve(l); err != nil {
		log.Println(err)
	}
}
