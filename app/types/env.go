package types

type Env struct {
	Port       string `env:"PORT"`
	DbURL      string `env:"DB_URL"`
	DbName     string `env:"DB_NAME"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
}
