package service

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mhrlife/centrifugo-chat-tutorial/config"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/ent"
	"github.com/stretchr/testify/require"
	"github.com/teris-io/shortid"
	"testing"
)

func NewTestingService(t *testing.T, config *config.Config) *Service {
	if config.App.Secret == "" {
		config.App.Secret = "secret"
	}

	fileID := shortid.MustGenerate()
	client, err := ent.Open("sqlite3", "file:"+fileID+"?mode=memory&cache=shared&_fk=1")

	require.NoError(t, err)
	require.NoError(t, client.Schema.Create(context.Background()))

	return NewService(client, config)
}
