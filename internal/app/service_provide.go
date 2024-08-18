package app

import (
	"context"
	"log"

	"github.com/ArtEmerged/o_chat-server/internal/api/grpc/chat"
	"github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/client/db/pg"
	"github.com/ArtEmerged/o_chat-server/internal/client/db/transaction"
	"github.com/ArtEmerged/o_chat-server/internal/closer"
	"github.com/ArtEmerged/o_chat-server/internal/config"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
	chatRepo "github.com/ArtEmerged/o_chat-server/internal/repository/chat"
	messageRepo "github.com/ArtEmerged/o_chat-server/internal/repository/message"
	"github.com/ArtEmerged/o_chat-server/internal/service"
	chatServ "github.com/ArtEmerged/o_chat-server/internal/service/chat"
	messageServ "github.com/ArtEmerged/o_chat-server/internal/service/message"
)

type serviceProvider struct {
	globalConfig *config.Config

	dbClient  db.Client
	txManager db.TxManager

	chatRepository    repository.ChatRepo
	messageRepository repository.MessageRepo

	chatService    service.ChatService
	messageService service.MessageService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GlobalConfig() *config.Config {
	if s.globalConfig == nil {
		s.globalConfig = config.New()

		err := s.globalConfig.Init("")
		if err != nil {
			log.Fatalf("failed init config: %v", err)
		}
	}

	return s.globalConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GlobalConfig().GetDbDNS())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %v", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepo {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.New(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepo {
	if s.messageRepository == nil {
		s.messageRepository = messageRepo.New(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatServ.New(s.ChatRepository(ctx), s.TxManager(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if s.messageService == nil {
		s.messageService = messageServ.New(s.MessageRepository(ctx), s.TxManager(ctx))
	}

	return s.messageService
}

func (s *serviceProvider) ChatImplementation(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx), s.MessageService(ctx))
	}

	return s.chatImpl
}
