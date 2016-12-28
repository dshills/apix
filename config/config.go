package config

import "os"

func env(prefix, name string) string {
	return os.Getenv(prefix + name)
}
