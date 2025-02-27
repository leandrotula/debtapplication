package token

type CreateTokenPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}

type Generator interface {
	GenerateJwtToken(secret, audience, issuer string) (string, error)
}

type CustomToken struct {
	secret   string
	audience string
	issuer   string
}

func (c *CustomToken) GenerateJwtToken(secret, audience, issuer string) (string, error) {
	return "", nil
}
