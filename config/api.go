package config

const (
	EnvSandbox    Env = "sandbox"
	EnvProduction Env = "production"
)

type Env string

var EnvURL = map[Env]string{
	EnvSandbox:    "https://tdev.kiriminaja.com",
	EnvProduction: "https://client.kiriminaja.com",
}
