package grpc

import (
	"context"
	"workflow_engine/internal/delivery/grpc/pb"
)

func (s *WorkflowGRPCServer) Signal(ctx context.Context, req *pb.SignalRequest) (*pb.SignalResponse, error) {
	signal, err := mapToDomainSignal(req)
	if err != nil {
		return &pb.SignalResponse{
			Status:      "ERROR",
			Description: err.Error(),
		}, nil
	}

	err = s.workflowUC.HandleSignal(
		ctx,
		req.WorkflowID,
		signal,
	)
	if err != nil {
		return &pb.SignalResponse{
			Status:      "ERROR",
			Description: err.Error(),
		}, nil
	}

	return &pb.SignalResponse{
		Status:      "OK",
		Description: "signal processed",
	}, nil
}
