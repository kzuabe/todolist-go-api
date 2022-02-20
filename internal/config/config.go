package config

import "os"

var API_ENV string
var DSN string

func init() {
	API_ENV = os.Getenv("API_ENV")
	DSN = os.Getenv("DSN")
}
