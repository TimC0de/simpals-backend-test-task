package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	variables map[string]string
}

func (env *Environment) Initialize(keys []string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env.variables = make(map[string]string)
	for _, key := range keys {
		val := os.Getenv(key)
		if val != "" {
			env.variables[key] = val
		}
	}
}

func (env *Environment) GetVal(key string) string {
	val, ok := env.variables[key]
	if !ok {
		log.Fatalf("Unable to read %s env variable\n", key)
		return ""
	}
	return val
}

func (env *Environment) GetValNoStrict(key string, defVal string) string {
	val, ok := env.variables[key]
	if !ok {
		return val
	}
	return defVal
}

func NewEnvironment(keys []string) *Environment {
	env := new(Environment)
	env.Initialize(keys)
	return env
}
