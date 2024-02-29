package env

import (
	envparser "github.com/caarlos0/env/v10"
)

type Env struct {
	FirebaseCredentialsFile string `env:"FIREBASE_CREDENTIALS_FILE" json:"firebase_credentials_file"`
}

func NewEnv() (*Env, error) {
	env := Env{}
	if err := envparser.Parse(&env); err != nil {
		return nil, err
	}

	return &env, nil
}
