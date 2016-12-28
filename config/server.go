package config

// Server is the config for the server
type Server struct {
	DocRoot string
	Env     string
	Host    string
	Log     string
	Port    string
	Prefix  string
	JWTKey  string
}

// NewServer will return the config for the server
func NewServer(prefix string) *Server {
	return &Server{
		DocRoot: env(prefix, "_SERVER_DOCROOT"),
		Env:     env(prefix, "_SERVER_ENV"),
		Host:    env(prefix, "_SERVER_HOST"),
		Log:     env(prefix, "_SERVER_LOG"),
		Port:    env(prefix, "_SERVER_PORT"),
		Prefix:  prefix,
		JWTKey:  env(prefix, "_JWT_KEY"),
	}
}
