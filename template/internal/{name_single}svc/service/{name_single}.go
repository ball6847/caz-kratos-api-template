package service

import (
	"context"
	"errors"

	pb "${go_module_name}/api/${namespace_single}/v1"
	"${go_module_name}/internal/${name_single}svc/biz"
	"${go_module_name}/internal/${name_single}svc/data"

	"github.com/go-kratos/kratos/v2/log"
)

// ${name_pascal}Service provide service for interprocess communication for ${name_lower} service
type ${name_pascal}Service struct {
	pb.Unimplemented${name_pascal}Server

	uc  *biz.${name_pascal}Usecase
	log *log.Helper
}

func New${name_pascal}Service(uc *biz.${name_pascal}Usecase, logger log.Logger) *${name_pascal}Service {
	return &${name_pascal}Service{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *${name_pascal}Service) Create${name_pascal}(ctx context.Context, req *pb.Create${name_pascal}Request) (*pb.Create${name_pascal}Reply, error) {
	return &pb.Create${name_pascal}Reply{}, nil
}

func (s *${name_pascal}Service) Update${name_pascal}(ctx context.Context, req *pb.Update${name_pascal}Request) (*pb.Update${name_pascal}Reply, error) {
	return &pb.Update${name_pascal}Reply{}, nil
}

func (s *${name_pascal}Service) Delete${name_pascal}(ctx context.Context, req *pb.Delete${name_pascal}Request) (*pb.Delete${name_pascal}Reply, error) {
	return &pb.Delete${name_pascal}Reply{}, nil
}

func (s *${name_pascal}Service) Get${name_pascal}(ctx context.Context, req *pb.Get${name_pascal}Request) (*pb.Get${name_pascal}Reply, error) {
	item, err := s.uc.Get${name_pascal}(ctx, req.Id)
	if err != nil {
		if errors.Is(err, data.ErrNoRows) {
			return nil, pb.Error${name_pascal}NotFound("${name_lower} not found")
		}
		s.log.WithContext(ctx).Error(err)
		return nil, pb.ErrorUnknown("error fetching ${name_lower}")
	}
	return &pb.Get${name_pascal}Reply{Data: make${name_pascal}Pb(item)}, nil
}

func (s *${name_pascal}Service) List${name_pascal}(ctx context.Context, req *pb.List${name_pascal}Request) (*pb.List${name_pascal}Reply, error) {
	items, err := s.uc.List${name_pascal}(ctx)
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, pb.ErrorUnknown("error fetching ${name_lower}")
	}
	return &pb.List${name_pascal}Reply{Data: make${name_pascal}PbList(items)}, nil
}
