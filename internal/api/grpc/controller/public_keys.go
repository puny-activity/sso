package controller

import (
	"context"
	"errors"
	"sso/internal/errs"
	"sso/pkg/contracts/ssocontract"
)

func (c *Controller) PublicKeys(ctx context.Context, request *ssocontract.PublicKeysRequest) (*ssocontract.PublicKeysResponse, error) {
	keys, err := c.keysUseCase.PublicKeys(ctx)
	if err != nil {
		c.log.Error(errs.Wrap("failed to get public keys", err).Error())
		return nil, errors.New(errs.ExtractText(err))
	}

	responseKeys := make([]string, len(keys))
	for i := range keys {
		responseKeys[i] = keys[i].String()
	}

	return &ssocontract.PublicKeysResponse{
		PublicKeys: responseKeys,
	}, nil
}
