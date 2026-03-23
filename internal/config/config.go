package config

import (
	"log"
	"sync"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	// Serve
	Port int    `env:"PORT" envDefault:"8081"`
	Env  string `env:"ENV" envDefault:"development"`

	// Database
	DbUrl    string `env:"DB_URL,required"`
	MaxConns int    `env:"MAX_CONNECTIONS" envDefault:"10"`
	MinConns int    `env:"MIN_CONNECTIONS" envDefault:"1"`

	// Business Logic
	VideoStoragePath string        `env:"VIDEO_STORAGE_PATH" envDefault:"./storage"`
	MaxUploadSize    int           `env:"MAX_UPLOAD_SIZE_MB" envDefault:"100"`
	ShutdownTimeout  time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		cfg = &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("Config error: %v", err)
		}
	})

	return cfg
}
