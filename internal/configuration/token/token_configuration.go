package token

import "time"

type ConfigurationToken struct {
	Secret     string
	Issuer     string
	Audience   string
	Expiration time.Duration
}

func NewConfigurationToken(expirationToken time.Duration, secret, issuer, audience string) *ConfigurationToken {
	return &ConfigurationToken{
		Secret:     secret,
		Audience:   audience,
		Issuer:     issuer,
		Expiration: expirationToken,
	}
}
