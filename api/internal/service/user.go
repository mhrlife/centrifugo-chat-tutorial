package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/apperror"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/ent"
)

type RegisterUser struct {
	JWT  string    `json:"jwt"`
	User *ent.User `json:"user"`
}

func (s *Service) RegisterUser(ctx context.Context, displayName string) (RegisterUser, error) {
	user, err := s.client.User.Create().
		SetDisplayName(displayName).
		Save(ctx)
	if err != nil {
		return RegisterUser{}, apperror.Wrap(err, "failed to create user")
	}

	token, err := s.generateJWT(user.ID.String())
	if err != nil {
		return RegisterUser{}, apperror.Wrap(err, "failed to generate JWT")
	}

	return RegisterUser{
		JWT:  token,
		User: user,
	}, nil
}

func (s *Service) generateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
	})

	tokenString, err := token.SignedString([]byte(s.config.App.Secret))
	if err != nil {
		return "", apperror.Wrap(err, "failed to sign token")
	}

	return tokenString, nil
}

func (s *Service) ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.App.Secret), nil
	})
	if err != nil {
		return "", apperror.Wrap(err, "failed to parse token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", apperror.NewUnauthorizedError("invalid token")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", apperror.NewUnauthorizedError("invalid token")
	}

	return userID, nil
}
