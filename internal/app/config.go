package app

import (
	"sso/internal/api/grpc"
	"sso/internal/infrastructure/keystorage"
	"sso/pkg/dbconnector"
)

type Config interface {
	Database() dbconnector.Config
	RefreshToken() keystorage.TokenConfig
	AccessToken() keystorage.TokenConfig
	GRPC() grpc.Config
}
