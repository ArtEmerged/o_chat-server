package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ArtEmerged/o_chat-server/config"
	grpc_chat "github.com/ArtEmerged/o_chat-server/internal/grpc/chat"
	repo "github.com/ArtEmerged/o_chat-server/internal/repository"
	serv "github.com/ArtEmerged/o_chat-server/internal/service"
	"github.com/ArtEmerged/o_chat-server/pkg/database"
)

func main() {
	cfg := config.New()
	err := cfg.Init("")
	if err != nil {
		log.Printf("failed init config: %v\n", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := database.NewPostgres(ctx, cfg.GetDbDNS())
	if err != nil {
		log.Printf("failed connection to postgres db: %v\n", err)
		return
	}
	defer db.Close()

	repository := repo.New(db)
	service := serv.New(repository)

	l, err := net.Listen("tcp", cfg.GetServerAddress())
	if err != nil {
		log.Printf("failed listen: %v\n", err)
		return
	}
	defer func() {
		err = l.Close()
		if err != nil {
			log.Printf("failed close listener: %v\n", err)
		}
	}()

	s := grpc.NewServer()
	reflection.Register(s)
	grpc_chat.Register(s, service)

	go func() {
		log.Printf("listen on :%s", cfg.ServerPort)
		if err = s.Serve(l); err != nil {
			log.Println(err)
		}
	}()

	quit, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-quit.Done()

	s.GracefulStop()
	log.Println("GRPC server has been shut down")
	db.Close()
	log.Println("Database closed")
}
