package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Env         string `mapstructure:"ENV"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // procura config.yaml no diretório raiz
	viper.AutomaticEnv()     // sobrescreve com variáveis de ambiente

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("ENV", "dev")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️ Config file não encontrado: %v (usando defaults/env)", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("❌ Erro ao carregar config: %v", err)
	}

	return &cfg
}
