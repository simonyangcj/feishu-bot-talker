package config

import "fmt"

type Config struct {
	AppID             string `json:"app_id" yaml:"appId"`
	AppSecret         string `json:"app_secret" yaml:"appSecret"`
	EncryptKey        string `json:"encrypt_key" yaml:"encryptKey"`
	VerificationToken string `json:"verification_token" yaml:"verificationToken"`
}

func (config *Config) validationConfig() error {
	if config.AppID == "" {
		return fmt.Errorf("app_id is required")
	}

	if config.AppSecret == "" {
		return fmt.Errorf("app_secret is required")
	}

	if config.EncryptKey == "" {
		return fmt.Errorf("encrypt_key is required")
	}

	if config.VerificationToken == "" {
		return fmt.Errorf("verification_token is required")
	}

	return nil
}
