package options

import (
	"os"
)

var API_KEY string
var ANNOUNCE_URL string
var ANTHELION_API_URL string

func LoadEnv() {
	API_KEY = os.Getenv("ANTHELION_API_KEY")
	ANNOUNCE_URL = os.Getenv("ANTHELION_ANNOUNCE_URL")
	ANTHELION_API_URL = os.Getenv("ANTHELION_API_URL")
}