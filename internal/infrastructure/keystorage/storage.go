package keystorage

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log/slog"
	"sso/internal/entity/publickey"
	"sso/internal/errs"
)

type Storage struct {
	refreshPrivateKey *rsa.PrivateKey
	accessPrivateKey  *rsa.PrivateKey
	log               *slog.Logger
}

func New(log *slog.Logger) *Storage {
	return &Storage{
		log: log,
	}
}

type TokenConfig interface {
	SecretKey() string
}

func (s *Storage) SetRefreshSecret(config TokenConfig) error {
	block, _ := pem.Decode([]byte(config.SecretKey()))
	if block == nil {
		return fmt.Errorf("failed to decode pem")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		s.refreshPrivateKey = privateKey
		return nil
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return errs.Wrap("failed to parse private key", err)
	}

	rsaPrivateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return fmt.Errorf("failed to parse private key")
	}
	s.refreshPrivateKey = rsaPrivateKey

	return nil
}

func (s *Storage) SetAccessSecret(config TokenConfig) error {
	block, _ := pem.Decode([]byte(config.SecretKey()))
	if block == nil {
		return fmt.Errorf("failed to decode pem")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		s.refreshPrivateKey = privateKey
		return nil
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return errs.Wrap("failed to parse private key", err)
	}

	rsaPrivateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return fmt.Errorf("failed to parse private key")
	}
	s.accessPrivateKey = rsaPrivateKey

	return nil
}

func (s *Storage) GetPublicKeys() ([]publickey.Key, error) {
	rsaPublicKey := &s.accessPrivateKey.PublicKey

	pubASN1, err := x509.MarshalPKIXPublicKey(rsaPublicKey)
	if err != nil {
		return nil, errs.Wrap("failed to marshal public key", err)
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	publicKey, err := publickey.New(string(pubPEM))
	if err != nil {
		return nil, errs.Wrap("failed to parse public key", err)
	}

	return []publickey.Key{publicKey}, nil
}
