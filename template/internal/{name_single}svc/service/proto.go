package service

import (
	v1 "${go_module_name}/api/${namespace_single}/v1"
	"${go_module_name}/internal/${name_single}svc/core/domain"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func make${name_pascal}Pb(data *domain.${name_pascal}) *v1.${name_pascal} {
	return &v1.${name_pascal}{
		Id: data.ID,
	}
}

func make${name_pascal}PbList(ww []*domain.${name_pascal}) []*v1.${name_pascal} {
	var items []*v1.${name_pascal}
	for _, w := range ww {
		items = append(items, make${name_pascal}Pb(w))
	}
	return items
}
