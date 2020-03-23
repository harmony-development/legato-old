package env

import "os"

var (
	// Verbosity is how "loud" the logging is for the server
	Verbosity string
	DBUser    string
	DBPass    string
	DBHost    string
	DBPort    string
)

// LoadVars reads environment variables into memory
func LoadVars() {
	Verbosity = os.Getenv("VERBOSITY_LEVEL")
	DBUser = os.Getenv("HARMONY_AUTH_USER")
	DBPass = os.Getenv("HARMONY_AUTH_PASSWORD")
	DBHost = os.Getenv("HARMONY_AUTH_HOST")
	DBPort = os.Getenv("HARMONY_AUTH_PORT")
}
