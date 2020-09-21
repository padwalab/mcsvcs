package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/padwalab/mcsvcs/pkg/custom"
)

type Set struct {
	TestStatusEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc custom.Service) Set {
	return Set{
		TestStatusEndpoint: MakeTestStatusEndpoint(svc),
	}
}

func MakeTestStatusEndpoint(svc custom.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(TestStatusRequest)
		code, err := svc.TestStatus(ctx)
		if err != nil {
			return TestStatusResponse{
				Code: code,
				Err:  err.Error(),
			}, nil
		}
		return TestStatusResponse{
			Code: code,
			Err:  "",
		}, nil
	}
}

func (s *Set) TestStatus(ctx context.Context) (int, error) {
	resp, err := s.TestStatusEndpoint(ctx, TestStatusRequest{})
	svcStatusResp := resp.(TestStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}
