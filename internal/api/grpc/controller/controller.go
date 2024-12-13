package controller

import (
	"context"
	"log/slog"
	"sso/internal/entity/publickey"
	"sso/pkg/contracts/ssocontract"
)

type Controller struct {
	ssocontract.SSOServer
	keysUseCase    keyUseCase
	accountUseCase accountUseCase
	tokenUseCase   tokenUseCase
	log            *slog.Logger
}

func New(keyUseCase keyUseCase, accountUseCase accountUseCase, tokenUseCase tokenUseCase,
	log *slog.Logger) *Controller {
	return &Controller{
		keysUseCase:    keyUseCase,
		accountUseCase: accountUseCase,
		tokenUseCase:   tokenUseCase,
		log:            log,
	}
}

type keyUseCase interface {
	PublicKeys(context.Context) ([]publickey.Key, error)
}

type accountUseCase interface {
	//Register(context.Context, *ssocontract.RegisterRequest) (*ssocontract.RegisterResponse, error)
	//Authenticate(context.Context, *ssocontract.AuthenticateRequest) (*ssocontract.AuthenticateResponse, error)
	//ChangePassword(context.Context, *ssocontract.ChangePasswordRequest) (*ssocontract.ChangePasswordResponse, error)
}

type tokenUseCase interface {
	//Refresh(context.Context, *ssocontract.RefreshRequest) (*ssocontract.RefreshResponse, error)
	//Revoke(context.Context, *ssocontract.RevokeRequest) (*ssocontract.RevokeResponse, error)
	//RevokeAll(context.Context, *ssocontract.RevokeAllRequest) (*ssocontract.RevokeAllResponse, error)
}
