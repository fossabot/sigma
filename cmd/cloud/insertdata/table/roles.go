package table

import (
	"log"
	"sigma/domain/core"
	r "sigma/domain/core/access/resource"
	"sigma/domain/managing/model"
	"sigma/domain/managing/repo"
	"sigma/domain/managing/service"
	"strings"
)

func InsertRoles(e *core.Engine) {
	roleRepo := repo.ProvideRepo(e)
	roleService := service.ProvideService(roleRepo)
	roles := []model.Role{
		{
			Name: "Super-Admin",
			Resources: strings.Join([]string{
				r.SupperAccess,
				r.CompanyRead, r.CompanyWrite, r.CompanyExcel,
				r.NodeRead, r.NodeWrite,
				r.UserNames, r.UserWrite, r.UserRead, r.UserReport,
				r.ActivitySelf, r.ActivityAll,
				r.AccountNames, r.AccountRead, r.AccountWrite, r.AccountExcel,
				r.RoleRead, r.RoleWrite,
			}, ", "),
			Description: "super-admin has all privileges - do not edit",
		},
		{
			Name: "Admin",
			Resources: strings.Join([]string{
				r.CompanyRead, r.CompanyWrite, r.CompanyExcel,
				r.UserNames, r.UserWrite, r.UserRead, r.UserReport,
				r.ActivitySelf, r.ActivityAll,
				r.AccountNames, r.AccountRead, r.AccountWrite, r.AccountExcel,
				r.RoleRead, r.RoleWrite,
			}, ", "),
			Description: "admin has all privileges - do not edit",
		},
		{
			Name:        "Cashier",
			Resources:   strings.Join([]string{r.ActivitySelf}, ", "),
			Description: "cashier has all privileges - after migration reset",
		},
	}
	roles[0].ID = 1
	roles[1].ID = 2
	roles[2].ID = 3

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			// e.ServerLog.Fatal(err)
			log.Fatalln(err.Error())
		}

	}

}
