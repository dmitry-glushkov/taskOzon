package server

type Config struct {
	Port    string `toml:"port"`
	MaxSize int    `toml:"max_size"`
}

func NewConfig() *Config {
	return &Config{
		Port:    ":8080",
		MaxSize: 100,
	}
}
