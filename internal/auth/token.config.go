package auth

type TokenConfig struct {
	Secret    []byte `env:"JWT_SECRET" envDefault:"secret"`
	ExpiredIn int    `env:"JWT_EXPIRED_IN" envDefault:"60"` // minute
}
