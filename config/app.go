package config

import "sso/pkg/dbconnector"

type App struct {
	database     Database
	refreshToken Token
	accessToken  Token
	grpc         GRPCServer
}

func (c App) Database() dbconnector.Config {
	return c.database
}

func (c App) GRPC() GRPCServer {
	return c.grpc
}
