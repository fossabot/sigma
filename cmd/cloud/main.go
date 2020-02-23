package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sigma/cmd/cloud/server"
)

func main() {

	s := server.Initialize()

	if err := s.Run(fmt.Sprintf("%v:%v", "localhost", 8183)); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Cloud Server started \n\n")
}
