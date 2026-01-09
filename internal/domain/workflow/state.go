package workflow

import "workflow_engine/internal/domain/entities"

var transitions = map[entities.DocumentStatus][]entities.DocumentStatus{
	entities.DocumentStatusStarted: {
		entities.DocumentStatusManagerConfirmed,
		entities.DocumentStatusManagerFailed,
	},
	entities.DocumentStatusManagerConfirmed: {
		entities.DocumentStatusEconomistConfirmed,
		entities.DocumentStatusEconomistFailed,
	},
	entities.DocumentStatusEconomistConfirmed: {
		entities.DocumentStatusBossConfirmed,
		entities.DocumentStatusBossFailed,
	},
	entities.DocumentStatusBossConfirmed: {
		entities.DocumentStatusSuccess,
	},
}

var rolePermissions = map[entities.UserRole][]entities.DocumentStatus{
	entities.RoleManager: {
		entities.DocumentStatusManagerConfirmed,
		entities.DocumentStatusManagerFailed,
	},
	entities.RoleEconomist: {
		entities.DocumentStatusEconomistConfirmed,
		entities.DocumentStatusEconomistFailed,
	},
	entities.RoleBoss: {
		entities.DocumentStatusBossConfirmed,
		entities.DocumentStatusBossFailed,
	},
}
