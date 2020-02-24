package service

import (
	"sigma/domain/core"
	"sigma/domain/managing/model"
	"sigma/domain/managing/repo"
)

// Service for injecting auth repo
type Service struct {
	Repo   repo.Repo
	Engine *core.Engine
}

// ProvideService for role is used in wire
func ProvideService(p repo.Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

// Save role
func (p *Service) Save(role model.Role) (createdRole model.Role, err error) {
	createdRole, err = p.Repo.Save(role)

	// p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving role for %+v", role))

	return
}
