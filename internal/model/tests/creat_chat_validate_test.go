package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

func TestCreateChatRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request model.CreateChatRequest
		wantErr error
	}{
		{
			name: "Valid request",
			request: model.CreateChatRequest{
				ChatName: "Foo",
			},
			wantErr: nil,
		},
		{
			name:    "Missing chat name",
			request: model.CreateChatRequest{},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field chat_name is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			require.Equal(t, tt.wantErr, err)
		})
	}
}
