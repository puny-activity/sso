package config

import (
	"sso/internal/api/grpc"
	"sso/internal/infrastructure/keystorage"
	"sso/pkg/dbconnector"
)

type App struct {
	database     Database
	refreshToken Token
	accessToken  Token
	grpc         GRPCServer
}

func (c App) Database() dbconnector.Config {
	return c.database
}

func (c App) RefreshToken() keystorage.TokenConfig {
	return c.refreshToken
}

func (c App) AccessToken() keystorage.TokenConfig {
	return c.accessToken
}

func (c App) GRPC() grpc.Config {
	return c.grpc
}
