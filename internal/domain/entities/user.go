package entities

import "strings"

type UserRole string

const (
	RoleWorker       UserRole = "worker"
	RoleManager      UserRole = "manager"
	RoleEconomist    UserRole = "economist"
	RoleBoss         UserRole = "boss"
	RoleInvestor     UserRole = "investor"
	RoleMainInvestor UserRole = "main_investor"
)

var validRoles = map[string]UserRole{
	"worker":        RoleWorker,
	"manager":       RoleManager,
	"economist":     RoleEconomist,
	"boss":          RoleBoss,
	"investor":      RoleInvestor,
	"main_investor": RoleMainInvestor,
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Phone     string
	Password  string
	Role      UserRole
}

func ParseUserRole(input string) (UserRole, error) {
	role, ok := validRoles[strings.ToLower(input)]
	if !ok {
		return "", ErrInvalidRole
	}
	return role, nil
}
