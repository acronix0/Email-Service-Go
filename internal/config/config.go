package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type Config struct {
	Env string `yaml:"env" env:"ENV" env-default:"local"`
	SMTPConfig SMTPConfig `yaml:"smtp_config"`
	OrderInfo OrderInfo `yaml:"order_info"`
	KafkaConfig KafkaConfig `yaml:"kafka_config"`
}
type SMTPConfig struct {
	Host           string `yaml:"host" env:"SMTP_HOST"`
	Username       string `yaml:"username" env:"SMTP_USERNAME"`
	Password       string `yaml:"password" env:"SMTP_PASSWORD"`
	Port           int    `yaml:"port" env:"SMTP_PORT" env-default:"587"`
}
type OrderInfo struct {
	Recipient string `yaml:"recipient"`
	Subject   string `yaml:"subject"`
	Template  string `yaml:"template"`
}

type KafkaConfig struct {
	BootstrapServer []string   `yaml:"bootstrap_servers"`
	Topics          []string `yaml:"topics""`
	ConsumerGroup string   `yaml:"consumer_group"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}
	return MustLoadPath(configPath)
}
func MustLoadPath(filePath string) *Config {
	if filePath == "" {
		log.Fatal("CONFIG_PATH not set")
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist\n", filePath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		log.Fatalf("cannot read config %s", filePath)
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}