package workflow

import "workflow_engine/internal/domain/entities"

func CanTransition(from, to entities.DocumentStatus) bool {
	allowedTransitions, exists := transitions[from]
	if !exists {
		return false
	}

	for _, allowed := range allowedTransitions {
		if allowed == to {
			return true
		}
	}
	return false
}

func CanRoleSetStatus(role entities.UserRole, status entities.DocumentStatus) bool {
	allowedStatuses, exists := rolePermissions[role]
	if !exists {
		return false
	}

	for _, allowed := range allowedStatuses {
		if allowed == status {
			return true
		}
	}
	return false
}

func GetNextStatus(current entities.DocumentStatus, role entities.UserRole) (entities.DocumentStatus, bool) {
	switch {
	case current == entities.DocumentStatusStarted && role == entities.RoleManager:
		return entities.DocumentStatusManagerConfirmed, true
	case current == entities.DocumentStatusManagerConfirmed && role == entities.RoleEconomist:
		return entities.DocumentStatusEconomistConfirmed, true
	case current == entities.DocumentStatusEconomistConfirmed && role == entities.RoleBoss:
		return entities.DocumentStatusBossConfirmed, true
	case current == entities.DocumentStatusBossConfirmed:
		return entities.DocumentStatusSuccess, true
	default:
		return "", false
	}
}

func GetFailedStatus(current entities.DocumentStatus, role entities.UserRole) (entities.DocumentStatus, bool) {
	switch {
	case current == entities.DocumentStatusStarted && role == entities.RoleManager:
		return entities.DocumentStatusManagerFailed, true
	case current == entities.DocumentStatusManagerConfirmed && role == entities.RoleEconomist:
		return entities.DocumentStatusEconomistFailed, true
	case current == entities.DocumentStatusEconomistConfirmed && role == entities.RoleBoss:
		return entities.DocumentStatusBossFailed, true
	default:
		return "", false
	}
}

func IsTerminalStatus(status entities.DocumentStatus) bool {
	return status == entities.DocumentStatusSuccess ||
		status == entities.DocumentStatusManagerFailed ||
		status == entities.DocumentStatusEconomistFailed ||
		status == entities.DocumentStatusBossFailed ||
		status == entities.DocumentStatusExpired
}

func GetExpectedRole(status entities.DocumentStatus) (entities.UserRole, bool) {
	switch status {
	case entities.DocumentStatusStarted:
		return entities.RoleManager, true
	case entities.DocumentStatusManagerConfirmed:
		return entities.RoleEconomist, true
	case entities.DocumentStatusEconomistConfirmed:
		return entities.RoleBoss, true
	default:
		return "", false
	}
}
