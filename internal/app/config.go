package app

import "sso/pkg/dbconnector"

type Config interface {
	Database() dbconnector.Config
}
