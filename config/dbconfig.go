package dbconfig

type Config struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}


func NewDBConfig(driver, username, password, host, dbName string, port int) *Config {
	return &Config{
		Driver:   driver,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbName,
	}
}