package config

import "net/http"

type Config struct {
	Env        Env
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}
