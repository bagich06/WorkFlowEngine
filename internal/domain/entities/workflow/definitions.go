package workflow

type WorkflowDefinition struct {
	Name  string
	Steps []WorkflowStep
}

type WorkflowStep struct {
	Name        string
	AllowedRole string
}

var DocumentApprovalWorkflow = WorkflowDefinition{
	Name: "document_approval",
	Steps: []WorkflowStep{
		{
			Name:        "manager_approval",
			AllowedRole: "manager",
		},
		{
			Name:        "economist_approval",
			AllowedRole: "economist",
		},
		{
			Name:        "boss_approval",
			AllowedRole: "boss",
		},
	},
}
