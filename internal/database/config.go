package database

type Config struct {
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://user:password@localhost/db"`
}
