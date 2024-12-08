package config

type Token struct {
	secretKey string
	ttlSecond int
}

func (c Token) SecretKey() string {
	return c.secretKey
}

func (c Token) TTLSecond() int {
	return c.ttlSecond
}
