package serializer

import "github.com/mhrlife/centrifugo-chat-tutorial/internal/ent"

type RegisterRequest struct {
	DisplayName string `json:"display_name" validator:"required,min=3,max=50"`
}

type UserInfo struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

func NewUserInfo(user *ent.User) UserInfo {
	return UserInfo{
		ID:          user.ID.String(),
		DisplayName: user.DisplayName,
	}
}

type UserWithToken struct {
	UserInfo

	Token string `json:"token"`
}

func NewUserWithToken(user *ent.User, token string) UserWithToken {
	return UserWithToken{
		UserInfo: NewUserInfo(user),
		Token:    token,
	}
}
