package grpc

import (
	pb "workflow_engine/internal/delivery/grpc/pb"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/entities/workflow"
)

func mapToDomainSignal(req *pb.SignalRequest) (workflow.WorkflowSignal, error) {
	var action workflow.WorkflowAction

	switch req.Action {
	case "approve":
		action = workflow.WorkflowActionApprove
	case "reject":
		action = workflow.WorkflowActionReject
	default:
		return workflow.WorkflowSignal{}, entities.ErrInvalidAction
	}

	role, err := entities.ParseUserRole(req.Role)
	if err != nil {
		return workflow.WorkflowSignal{}, err
	}

	return workflow.WorkflowSignal{
		Action: action,
		Role:   role,
	}, nil
}
