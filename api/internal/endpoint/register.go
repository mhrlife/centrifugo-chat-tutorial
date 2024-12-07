package endpoint

import (
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/apperror"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/serializer"
)

func (e *Endpoint) register(c *Context) error {
	request, err := BindAndValidate[serializer.RegisterRequest](c)
	if err != nil {
		return apperror.Wrap(err, "couldn't validate register request")
	}

	registeredUser, err := e.service.RegisterUser(c.Context(), request.DisplayName)
	if err != nil {
		return err
	}

	return c.Ok(serializer.NewUserWithToken(registeredUser.User, registeredUser.JWT))
}
