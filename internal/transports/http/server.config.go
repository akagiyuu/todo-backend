package http

type Config struct {
	Port       int    `env:"PORT" envDefault:"3000"`
	CorsOrigin string `env:"CORS_ORIGIN" envDefault:"*"`
}
