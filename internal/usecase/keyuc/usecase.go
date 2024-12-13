package keyuc

import (
	"log/slog"
	"sso/internal/entity/publickey"
)

type UseCase struct {
	keyStorage keyStorage
	log        *slog.Logger
}

func New(keyStorage keyStorage, log *slog.Logger) *UseCase {
	return &UseCase{
		keyStorage: keyStorage,
		log:        log,
	}
}

type keyStorage interface {
	GetPublicKeys() ([]publickey.Key, error)
}
