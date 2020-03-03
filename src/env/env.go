package env

import "os"

func EnvVarialbe(key string) string {
	// Set env variable using os package
	os.Setenv(key, "gopher")

	// return the env variable using os package
	return os.Getenv(key)
}
