package port

import (
	"${go_module_name}/internal/${name_single}svc/core/domain"
	"context"
)

type ${name_pascal}Repository interface {
	FindAll(context.Context) ([]*domain.${name_pascal}, error)
	FindOne(context.Context, string) (*domain.${name_pascal}, error)
}
