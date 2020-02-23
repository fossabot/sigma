package main

import (
	"sigma/cmd/cloud/server"
	"sigma/cmd/cloud/setup"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	engine := setup.LoadEnv()
	setup.InitDB(engine)
	setup.Migrate(engine)

	server.Initialize(engine)

}
