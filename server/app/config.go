package app

type config struct {
	ParseConfig    *ParseConfig
	DatabaseConfig *DatabaseConfig
}

type ParseConfig struct {
	ParseResumeURL string
	ParseAPIKey    string
}

type DatabaseConfig struct {
	JdbcUrl  string
	Username string
	Password string
}
