package credential

import "sso/internal/entity/credential/username"

type Uhashed struct {
	Username username.Username
}

type Hashed struct {
	Username username.Username
	Password string
}
