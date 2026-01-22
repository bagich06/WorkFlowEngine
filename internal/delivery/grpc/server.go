package grpc

import (
	"workflow_engine/internal/delivery/grpc/pb"
	"workflow_engine/internal/domain/usecase"
)

type WorkflowGRPCServer struct {
	pb.UnimplementedWorkflowServer
	workflowUC *usecase.WorkFlowUseCase
}

func NewWorkflowGRPCServer(
	workflowUC *usecase.WorkFlowUseCase,
) *WorkflowGRPCServer {
	return &WorkflowGRPCServer{
		workflowUC: workflowUC,
	}
}
