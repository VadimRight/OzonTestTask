package config

import (
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

type Config struct {
	Env *EnvConfig
	Postgres *PostgresConfig
	Server *ServerConfig
}

type EnvConfig struct {
	Env string 
	EnvPath string
}

type PostgresConfig struct {	
	PostgresPort string 
	PostgresHost string 
	DatabaseName string 
	PostgresUser string 
	PostgresPassword string 
}

type ServerConfig struct {
	ServerAddress string 
	ServerPort string 
	Timeout           time.Duration 
	IdleTimeout       time.Duration 
	RunMode string
}

func LoadConfig() *Config {
	envConfig := LoadEnvConfig()
	postgresConfig := LoadPostgresConfig()
	serverConfig := LoadServerConfig()	
	return &Config {
		Env: envConfig,
		Postgres: postgresConfig,
		Server: serverConfig,
	}	
}

func LoadEnvConfig() *EnvConfig {
	const opt = "internal.config.LoadEnvConfig"
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s %v", opt, err)
	}		

	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	configPath := os.Getenv("CONFIG_PATH")
	log.Printf("CONFIG_PATH is %s", configPath)
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exists", configPath)
	}

	envType, ok := os.LookupEnv("ENV")
	if !ok {
		log.Fatal("Can't read ENV")
	}
	return &EnvConfig {
		Env: envType, 
		EnvPath: configPath,
	}
}

func LoadPostgresConfig() *PostgresConfig {
	const opt = "internal.config.LoadPostgresConfig"
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s: %v", opt, err)
	}
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	
	postgresPort, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {log.Fatal("Can't read POSTGRES_PORT")}

	postgresHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {log.Fatal("Can't read POSTGRES_HOST")}

	postgresPassword, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {log.Fatal("Can't read POSTGRES_PASSWORD")}

	postgresDB, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {log.Fatal("Can't read POSTGRES_DB")}
	
	postgresUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {log.Fatal("Can't read POSTGRES_USER")}
	
	return &PostgresConfig {
		PostgresPort: postgresPort,
		PostgresHost: postgresHost,
		DatabaseName: postgresDB,
		PostgresUser: postgresUser,
		PostgresPassword: postgresPassword,	
	}
}

func LoadServerConfig() *ServerConfig {
	err := godotenv.Load()
	const opt = "internal.config.LoadPostgresConfig"
	if err != nil {
		log.Fatalf("%s: %v", opt, err)
	}
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	
	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if !ok {log.Fatal("Can't read SERVER_PORT")}
	
	serverAddr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {log.Fatal("Can't read SERVER_ADDR")}
	
	timeout, ok := os.LookupEnv("TIMEOUT")
	if !ok {log.Fatal("Can't read TIMEOUT")}
	
	timeoutTime, err := time.ParseDuration(timeout)
	if err != nil {log.Fatalf("error while parsing timeout")}
	
	idleTimeout, ok := os.LookupEnv("IDLE_TIMEOUT")
	if !ok {log.Fatal("Can't read IDLE_TIMEOUT")}
	
	idleTimeoutTime, err := time.ParseDuration(idleTimeout)
	if err != nil {	log.Fatalf("error while parsing idle time")}
	
	serverRunMode, ok := os.LookupEnv("SERVER_RUN_MODE")
	if !ok{	log.Fatalf("err while parsing run mode")}
	
	return &ServerConfig {
		ServerAddress: serverAddr, 
		ServerPort: serverPort,
		Timeout: timeoutTime, 
		IdleTimeout: idleTimeoutTime,  	
		RunMode: serverRunMode,
	}
}
