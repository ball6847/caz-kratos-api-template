package ${name_single}svc

import (
	"${go_module_name}/internal/${name_single}svc/biz"
	"${go_module_name}/internal/${name_single}svc/data"
	"${go_module_name}/internal/${name_single}svc/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	data.New${name_pascal}Repository,
	biz.New${name_pascal}Usercase,
	service.New${name_pascal}Service,
)
