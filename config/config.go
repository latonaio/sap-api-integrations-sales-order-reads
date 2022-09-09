package config

import (
	"fmt"
	"os"
	"strconv"
)

type Conf struct {
	SAP *SAP
}

func NewConf() *Conf {
	return &Conf{
		SAP: newSAP(),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		val = fallback
	}
	return val
}

func getEnvInt(key string, fallback int) int {
	rawVal := os.Getenv(key)
	val, err := strconv.Atoi(rawVal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "environment %s required number type: %+v", key, err)
		val = fallback
	}
	return val
}
