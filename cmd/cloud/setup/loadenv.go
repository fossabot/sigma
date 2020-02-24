package setup

import (
	envEngine "github.com/caarlos0/env/v6"
	"log"
	"sigma/domain/core"
)

// LoadEnv get variables from environment and put them inside engine.Env
func LoadEnv() *core.Engine {
	var engine core.Engine
	if err := envEngine.Parse(&engine.Env); err != nil {
		log.Fatalln(err)
	}

	return &engine
}
