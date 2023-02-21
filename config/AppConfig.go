package config

type AppConfig struct {
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     string
	DatabaseName     string
	DatabaseHost     string
	AppPort          string
	SecretKey        string
}
