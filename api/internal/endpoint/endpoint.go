package endpoint

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mhrlife/centrifugo-chat-tutorial/config"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/apperror"
	"github.com/mhrlife/centrifugo-chat-tutorial/internal/service"
	"github.com/sirupsen/logrus"
)

type Endpoint struct {
	config   *config.Config
	e        *echo.Echo
	validate *validator.Validate
	service  *service.Service
}

func NewEndpoint(cfg *config.Config, e *echo.Echo, svc *service.Service) *Endpoint {
	em := &Endpoint{
		e:        e,
		validate: validator.New(),
		service:  svc,
		config:   cfg,
	}

	em.init()

	return em
}

func (e *Endpoint) init() {
	auth := e.e.Group("auth")
	auth.POST("/register", e.wrapContext(e.register))

	e.e.HTTPErrorHandler = func(err error, c echo.Context) {
		ctx := NewContext(e, c)

		errCode := apperror.ExtErrorCode(err)
		errMessage := apperror.ExtErrorMessage(err)

		logrus.WithError(err).Error("error happened")

		_ = ctx.Error(errCode, errMessage)
	}
}

type Context struct {
	echo     echo.Context
	validate *validator.Validate
}

func NewContext(e *Endpoint, c echo.Context) *Context {
	return &Context{
		echo:     c,
		validate: e.validate,
	}

}

func (c *Context) Ok(data any) error {
	return c.echo.JSON(200, map[string]any{
		"ok":   true,
		"data": data,
	})
}

func (c *Context) Error(code int, message string) error {
	return c.echo.JSON(200, map[string]any{
		"ok":      false,
		"message": message,
		"code":    code,
	})
}

func (c *Context) Context() context.Context {
	return c.echo.Request().Context()
}

func (e *Endpoint) wrapContext(f func(ctx *Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := NewContext(e, c)

		return f(ctx)
	}
}

func (e *Endpoint) Start() {
	e.e.Logger.Error(e.e.Start(":" + e.config.App.Port))
}
