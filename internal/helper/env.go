package helper

import (
	"os"
	"strconv"
)

func GetEnvString(env string) string {
	return os.Getenv(env)
}

func GetEnvInt(env string) int {
	envStr := os.Getenv(env)

	parseEnv, _ := strconv.Atoi(envStr)

	return parseEnv
}
