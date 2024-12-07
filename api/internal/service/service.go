package service

import (
	"github.com/mhrlife/centrifugo-chat-tutorial/config"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/ent"
)

type Service struct {
	client *ent.Client
	config *config.Config
}

func NewService(client *ent.Client, config *config.Config) *Service {
	return &Service{
		client: client,
		config: config,
	}
}

func (s *Service) Close() {
	s.client.Close()
}
