package config

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetConfig() Config {
	return Config{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "anton132",
		DBName:   "achievement-holder",
		SSLMode:  "disable",
	}
}
