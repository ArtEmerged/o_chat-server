package app

import (
	"context"
	"log"

	"github.com/ArtEmerged/library/client/cache"
	"github.com/ArtEmerged/library/client/cache/redis"
	"github.com/ArtEmerged/library/client/db"
	"github.com/ArtEmerged/library/client/db/pg"
	"github.com/ArtEmerged/library/client/db/transaction"

	"github.com/ArtEmerged/o_chat-server/internal/api/grpc/chat"
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
	redisConfig  redis.Config

	dbClient  db.Client
	txManager db.TxManager
	cache     cache.Cache

	chatRepository    repository.ChatRepo
	messageRepository repository.MessageRepo

	chatService    service.ChatService
	messageService service.MessageService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GlobalConfig returns the global config.
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

// RedisConfig returns the redis config.
func (s *serviceProvider) RedisConfig() redis.Config {
	if s.redisConfig == nil {
		s.redisConfig = s.GlobalConfig().RedisConfig()
	}

	return s.redisConfig
}

// DBClient returns the db client.
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GlobalConfig().DbDNS())
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

// TxManager returns the transaction manager.
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// Cache returns the cache.
func (s *serviceProvider) Cache(ctx context.Context) cache.Cache {
	if s.cache == nil {
		s.cache = redis.NewClient(s.RedisConfig())
		err := s.cache.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping redis: %v", err)
		}

		closer.Add(s.cache.Close)
	}

	return s.cache
}

// ChatRepository returns the chat repository.
func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepo {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.New(s.DBClient(ctx), s.Cache(ctx))
	}

	return s.chatRepository
}

// MessageRepository returns the message repository.
func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepo {
	if s.messageRepository == nil {
		s.messageRepository = messageRepo.New(s.DBClient(ctx), s.Cache(ctx))
	}

	return s.messageRepository
}

// ChatService returns the chat service.
func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatServ.New(s.ChatRepository(ctx), s.TxManager(ctx), s.Cache(ctx))
	}

	return s.chatService
}

// MessageService returns the message service.
func (s *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if s.messageService == nil {
		s.messageService = messageServ.New(s.MessageRepository(ctx), s.Cache(ctx), s.TxManager(ctx))
	}

	return s.messageService
}

// ChatImplementation returns the chat implementation.
func (s *serviceProvider) ChatImplementation(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx), s.MessageService(ctx))
	}

	return s.chatImpl
}
