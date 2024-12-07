package service

import (
	"context"
	"fmt"
	"github.com/mhrlife/centrifugo-chat-tutorial/config"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/apperror"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite

	*Service
	ctx context.Context
}

func TestUserSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(UserTestSuite))
}

func (t *UserTestSuite) SetupSuite() {
	t.Service = NewTestingService(t.T(), &config.Config{})
	t.ctx = context.Background()
}

func (t *UserTestSuite) TearDownSuite() {
	t.Service.Close()
}

func (t *UserTestSuite) TestRegisterUser() {
	user, err := t.RegisterUser(t.ctx, "Ali")
	t.NoError(err)
	t.NotNil(user.User)
	t.NotEmpty(user.JWT)

	// Check if the user is saved in the database
	_, err = t.client.User.Get(t.ctx, user.User.ID)
	t.NoError(err)

	// Check if the JWT is valid
	userID, err := t.ValidateJWT(user.JWT)
	t.NoError(err)
	t.Equal(user.User.ID.String(), userID)
}

func (t *UserTestSuite) TestRegisterUserError() {
	_, err := t.RegisterUser(t.ctx, "")
	fmt.Println(err)
	t.Require().True(apperror.IsValidationError(err))
}
