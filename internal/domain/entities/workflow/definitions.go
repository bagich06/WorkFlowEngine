package workflow

import "workflow_engine/internal/domain/entities"

type WorkflowDefinition struct {
	Name  string
	Steps []WorkflowStep
}

type WorkflowContext struct {
	Amount float64
}

type WorkflowStep struct {
	Name          string
	AllowedRole   entities.UserRole
	ParallelGroup WorkflowGroup
}

var DocumentApprovalWorkflow = WorkflowDefinition{
	Name: "document_approval",
	Steps: []WorkflowStep{
		{
			Name:          "manager_approval",
			AllowedRole:   entities.RoleManager,
			ParallelGroup: WorkflowGroupFirst,
		},
		{
			Name:          "economist_approval",
			AllowedRole:   entities.RoleEconomist,
			ParallelGroup: WorkflowGroupFirst,
		},
		{
			Name:          "boss_approval",
			AllowedRole:   entities.RoleBoss,
			ParallelGroup: WorkflowGroupFirst,
		},
		{
			Name:          "investors_approval",
			AllowedRole:   entities.RoleInvestor,
			ParallelGroup: WorkflowGroupFinal,
		},
		{
			Name:          "final_approval",
			AllowedRole:   entities.RoleMainInvestor,
			ParallelGroup: WorkflowGroupFinal,
		},
	},
}
