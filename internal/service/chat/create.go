package chat

import (
	"context"
	"log"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

// CreateChat creates new chat by chat name and creator id with user ids.
func (s *chatService) CreateChat(ctx context.Context, in *model.CreateChatRequest) (id int64, err error) {

	if err = in.Validate(); err != nil {
		return -1, err
	}

	// add the chat owner to the list of participants
	in.UserIDs = append(in.UserIDs, in.CreatorID)

	// remove duplicates
	in.UserIDs = uniqueSliceInt64(in.UserIDs)

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		id, errTx = s.repo.CreateChat(ctx, in)
		if errTx != nil {
			return errTx
		}

		errTx = s.repo.AddUsersToChat(ctx, id, in.UserIDs)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return -1, err
	}

	newChat := &model.Chat{
		ID:        id,
		ChatName:  in.ChatName,
		CreatorID: in.CreatorID,
		UserIDs:   in.UserIDs,
	}

	err = s.cache.Set(ctx, model.ChatCacheKey(id), newChat, 0)
	if err != nil {
		log.Printf("WARN: failed to set chat in cache: %s\n", err.Error())
	}

	return id, nil
}
