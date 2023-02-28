/*
 * @FilePath: /ChatGPT-CLI/services/key.go
 * @Author: maggot-code
 * @Date: 2023-02-28 21:45:25
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-03-01 00:10:28
 * @Description:
 */
package services

import "github.com/zalando/go-keyring"

const (
	serviceName = "gpt-cli"
)

type KeyService struct {
	ServiceName string
}

type KeyInterface interface {
	SetKey(key string) error
	GetKey() (string, error)
	ClearKey() error
}

func UseKeyService() KeyInterface {
	return DefineServiceKey()
}

func DefineServiceKey() *KeyService {
	return &KeyService{
		ServiceName: serviceName,
	}
}

func (k *KeyService) SetKey(key string) error {
	keyring.Set(k.ServiceName, serviceName, key)
	return nil
}

func (k *KeyService) GetKey() (string, error) {
	get, err := keyring.Get(k.ServiceName, serviceName)
	if err != nil {
		return "", err
	}

	return get, nil
}

func (k *KeyService) ClearKey() error {
	keyring.Delete(k.ServiceName, serviceName)
	return nil
}
