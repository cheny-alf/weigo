package svc

import (
	"wego/apps/service/user/api/internal/config"
	"wego/apps/service/user/model/ent"
)

type ServiceContext struct {
	Config config.Config
	Client *ent.Client
}

func NewServiceContext(c config.Config, client *ent.Client) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Client: client,
	}
}
