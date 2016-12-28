package config

// SMTP is the config for connecting to the database
type SMTP struct {
	Email    string
	Host     string
	Name     string
	Password string
	Port     string
	Prefix   string
	User     string
}

// NewSMTP will return the config for the SMTP server
func NewSMTP(prefix string) *SMTP {
	return &SMTP{
		Host:     env(prefix, "_SMTP_HOST"),
		Password: env(prefix, "_SMTP_PASS"),
		Port:     env(prefix, "_SMTP_PORT"),
		Prefix:   prefix,
		User:     env(prefix, "_SMTP_USER"),
		Name:     env(prefix, "_SMTP_NAME"),
		Email:    env(prefix, "_SMTP_EMAIL"),
	}
}
