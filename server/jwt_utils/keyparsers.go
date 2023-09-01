package jwt_utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ParsePrivateKey() (*ecdsa.PrivateKey, error) {
	pemData, err := LoadPEMFile(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		return nil, fmt.Errorf("failed to load PEM file: %w", err)
	}

	block, _ := pem.Decode(pemData)

	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)

	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return key, nil
}

func ParsePublicKey() (*ecdsa.PublicKey, error) {
	pemData, err := LoadPEMFile(os.Getenv("PUBLIC_KEY"))

	if err != nil {
		return nil, fmt.Errorf("failed to load PEM file: %w", err)
	}

	block, _ := pem.Decode(pemData)

	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	publicKey := key.(*ecdsa.PublicKey)

	return publicKey, nil
}
