package insertdata

import (
	"fmt"
	"sigma/cmd/cloud/insertdata/table"
	"sigma/domain/core"
)

// Insert is used for add static rows to database
func Insert(engine *core.Engine) error {
	fmt.Println("KKKKKKKKK")

	table.InsertRoles(engine)

	return nil

}
