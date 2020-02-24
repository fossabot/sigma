package repo

import (
	"sigma/domain/core"
	"sigma/domain/managing/model"
)

// Repo for injecting engine
type Repo struct {
	Engine *core.Engine
}

// ProvideRepo is used in wire
func ProvideRepo(engine *core.Engine) Repo {
	return Repo{Engine: engine}
}

// Save role
func (p *Repo) Save(role model.Role) (u model.Role, err error) {
	err = p.Engine.DB.Save(&role).Scan(&u).Error
	return
}
