package config

type Config struct {
	app    App
	logger Logger
}

func Parse() (*Config, error) {
	return &Config{
		app: App{
			database: Database{
				database:       "postgresql",
				host:           "localhost",
				port:           4000,
				name:           "name",
				user:           "user",
				password:       "password",
				migrationsPath: "migrations/postgresql",
			},
			refreshToken: Token{
				secretKey: "refresh_secret_key",
				ttlSecond: 2592000,
			},
			accessToken: Token{
				secretKey: "access_secret_key",
				ttlSecond: 1800,
			},
			grpc: GRPCServer{
				port: 6000,
			},
		},
		logger: Logger{
			level: "debug",
		},
	}, nil
}

func (c Config) App() App {
	return c.app
}

func (c Config) Logger() Logger {
	return c.logger
}
