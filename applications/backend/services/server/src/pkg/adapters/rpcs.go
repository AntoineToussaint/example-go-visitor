package adapters

import (
	"context"
	"github.com/codefly-dev/go-grpc/base/pkg/business"
	"github.com/codefly-dev/go-grpc/base/pkg/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
Include your RPCs here
*/

var service *business.Service

func SetService(s *business.Service) {
	service = s
}

func (s *GrpcServer) CreateVisit(ctx context.Context, req *gen.CreateVisitRequest) (*gen.CreateVisitResponse, error) {
	if err := Validate(req); err != nil {
		return nil, err
	}
	number, err := service.Visit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gen.CreateVisitResponse{
		VisitNumber: int64(number),
	}, nil
}

func (s *GrpcServer) GetVisitStatistics(ctx context.Context, req *gen.GetVisitStatisticsRequest) (*gen.GetVisitStatisticsResponse, error) {
	if err := Validate(req); err != nil {
		return nil, err
	}
	stats, err := service.GetStatistics(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gen.GetVisitStatisticsResponse{
		TotalVisits: int64(stats.Total),
		Visits:      stats.Visits,
	}, nil
}
