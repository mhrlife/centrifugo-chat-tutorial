package endpoint

import "github.com/mhrlife/centrifugo-chat-tutorial/internal/apperror"

func BindAndValidate[T any](c *Context) (T, error) {
	var t T
	if err := c.echo.Bind(&t); err != nil {
		return t, apperror.NewValidationError("invalid request body").Wrap(err)
	}

	if err := c.validate.Struct(t); err != nil {
		return t, apperror.NewValidationError("validation failed").Wrap(err)
	}

	return t, nil
}
