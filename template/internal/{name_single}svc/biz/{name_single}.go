package biz

import (
	"${go_module_name}/internal/${name_single}svc/core/domain"
	"${go_module_name}/internal/${name_single}svc/core/port"
	"${go_module_name}/internal/${name_single}svc/data"
	"context"
)

type ${name_pascal}Usecase struct {
	repo port.${name_pascal}Repository
}

func New${name_pascal}Usercase(repo *data.${name_pascal}Repo) *${name_pascal}Usecase {
	return &${name_pascal}Usecase{repo: repo}
}

func (u *${name_pascal}Usecase) Get${name_pascal}(ctx context.Context, id string) (*domain.${name_pascal}, error) {
  return nil, nil
}

func (u *${name_pascal}Usecase) List${name_pascal}(ctx context.Context) ([]*domain.${name_pascal}, error) {
	return nil, nil
}
