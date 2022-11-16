package config

func GetConfig() *Config {
	return &Config{
		Server:   getServerConfig(),
		DataBase: getDataBaseConfig(),
	}
}

func getDataBaseConfig() DataBase { // database configuration
	return DataBase{
		DbHost:   "localhost",
		DbPort:   "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "postgres",
		SSLMode:  "disable",
	}
}

func getServerConfig() Server {
	return Server{
		Host: "localhost",
		Port: ":8080",
	}
}
