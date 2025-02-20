package data

import (
	"context"
	"${go_module_name}/internal/${name_single}svc/core/domain"
)

type ${name_pascal}Repository struct {
}

func New${name_pascal}Repository() *${name_pascal}Repository {
	return &${name_pascal}Repository{}
}

func (r *${name_pascal}Repository) FindAll(ctx context.Context) ([]*domain.${name_pascal}, error) {
  return nil, nil
}

func (r *${name_pascal}Repository) FindOne(ctx context.Context, id string) (*domain.${name_pascal}, error) {
  return nil, nil
}
