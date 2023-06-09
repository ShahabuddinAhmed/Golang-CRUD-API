package main

type Config struct {
	DBConfig DBConfig
	// Other configurations can be added here
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func LoadConfig() *Config {
	return &Config{
		DBConfig: DBConfig{
			Host:     "0.0.0.0",
			Port:     "3306",
			Username: "root",
			Password: "root",
			DBName:   "crud_demo",
		},
		// Other configurations can be added here
	}
}
