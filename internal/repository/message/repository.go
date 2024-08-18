package chat

import (
	dbClient "github.com/ArtEmerged/o_chat-server/internal/client/db"
	"github.com/ArtEmerged/o_chat-server/internal/repository"
)

const (
	tableChats = "public.chats"

	tableChatsIDColumn        = "id"
	tableChatsNameColumn      = "name"
	tableChatsOwnerColumn     = "owner"
	tableChatsCreatedAtColumn = "created_at"
	tableChatsDeletedAtColumn = "deleted_at"
)

const (
	tableChatUsers = "public.chat_users"

	tableChatUsersChatIDColumn = "chat_id"
	tableChatUsersUserIDColumn = "user_id"
)

const (
	tableMessages                 = "public.chat_messages"
	tableMessagesChatIDColumn     = "chat_id"
	tableMessagesFromUserIDColumn = "from_user_id"
	tableMessagesTextColumn       = "text"
	tableMessagesCreatedAtColumn  = "created_at"
)

type messageRepo struct {
	db dbClient.Client
}

// New creates a new instance of messageRepo with the given database connection pool.
// db - database connection pool
func New(db dbClient.Client) repository.MessageRepo {
	return &messageRepo{db: db}
}
