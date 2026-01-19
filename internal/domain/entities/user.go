package entities

type UserRole string

const (
	RoleWorker    UserRole = "worker"
	RoleManager   UserRole = "manager"
	RoleEconomist UserRole = "economist"
	RoleBoss      UserRole = "boss"
	RoleInvestor  UserRole = "investor"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Phone     string
	Password  string
	Role      UserRole
}
