package config

import (
	"fmt"
)

type Environment string

const (
	EnvDevelop    Environment = "development"
	EnvProduction Environment = "production"
)

func EnvironmentFromStr(s string) (Environment, error) {
	switch s {
	case string(EnvDevelop):
		return EnvDevelop, nil
	case string(EnvProduction):
		return EnvProduction, nil
	default:
		return "", fmt.Errorf("unknown environment: %s", s)
	}
}