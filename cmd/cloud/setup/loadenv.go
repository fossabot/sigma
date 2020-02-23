package setup

import (
	envEngine "github.com/caarlos0/env/v6"
	"log"
	"sigma/domain/core/engine"
)

// LoadEnv get variables from environment and put them inside engine.Env
func LoadEnv() *engine.Engine {
	var engine engine.Engine
	if err := envEngine.Parse(&engine.Env); err != nil {
		log.Fatalln(err)
	}

	return &engine
}
