package config

type config struct {
	ParseConfig    *ParseConfig
	DatabaseConfig *DatabaseConfig
}

var appConfig *config = &config{}

//GetConfig Get Config Singleton
func GetConfig() *config {
	return appConfig
}

type ParseConfig struct {
	ParseResumeURL string
	ParseAPIKey    string
}

type DatabaseConfig struct {
	JdbcUrl string
	Driver  string
}
