package config

type Configuration struct {
	Database Database
}

type Database struct {
	Host     string
	Port     int
	Dbname   string
	Username string
	Password string
}

var Config Configuration
