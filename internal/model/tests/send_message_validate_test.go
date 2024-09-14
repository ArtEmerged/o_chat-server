package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_chat-server/internal/model"
)

func TestSendMessageRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request model.SendMessageRequest
		wantErr error
	}{
		{
			name: "Valid request",
			request: model.SendMessageRequest{
				Text: "Hello, World!",
			},
			wantErr: nil,
		},
		{
			name:    "Missing text",
			request: model.SendMessageRequest{},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field text is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			require.Equal(t, tt.wantErr, err)
		})
	}
}
