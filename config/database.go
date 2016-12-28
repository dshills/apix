package config

// Database is the config for connecting to the database
type Database struct {
	Host     string
	Name     string
	Params   string
	Password string
	Port     string
	Prefix   string
	User     string
}

// NewDatabase will return the config for the database
func NewDatabase(prefix string) *Database {
	return &Database{
		Host:     env(prefix, "_DB_HOST"),
		Name:     env(prefix, "_DB_NAME"),
		Params:   env(prefix, "_DB_PARAMS"),
		Password: env(prefix, "_DB_PASS"),
		Port:     env(prefix, "_DB_PORT"),
		Prefix:   prefix,
		User:     env(prefix, "_DB_USER"),
	}
}
