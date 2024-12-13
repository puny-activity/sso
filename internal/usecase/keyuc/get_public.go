package keyuc

import (
	"context"
	"sso/internal/entity/publickey"
	"sso/internal/errs"
)

func (u *UseCase) PublicKeys(context.Context) ([]publickey.Key, error) {
	publicKeys, err := u.keyStorage.GetPublicKeys()
	if err != nil {
		return nil, errs.Wrap("failed to fetch public keys", err)
	}
	return publicKeys, nil
}
