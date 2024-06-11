// Package keystore implements a simple way to store & retreive secrets
package keystore

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

func StoreSecretIfNotExists(id, secret string) error {
	_, err := keyring.Get(buildServiceName(id), "")
	if err == nil {
		return fmt.Errorf("secret with ID %s already exists in keystore", id)
	}

	if err := keyring.Set(buildServiceName(id), "", secret); err != nil {
		return fmt.Errorf("unable to store secret in keystore: %v", err)
	}
	return nil
}

func RemoveSecret(id string) error {
	err := keyring.Delete(buildServiceName(id), "")
	if err != nil {
		return fmt.Errorf("unable to delete secret: %v", err)
	}
	return nil
}

func RetrieveSecret(id string) (string, error) {
	secret, err := keyring.Get(buildServiceName(id), "")
	if err != nil {
		return "", fmt.Errorf("unable to retrieve secret: %v", err)
	}
	return secret, nil
}

func buildServiceName(n string) string {
	return "TOTPer " + n
}
