package entities

type DocumentStatus string

const (
	DocumentStatusStarted  DocumentStatus = "started"
	DocumentStatusApproved DocumentStatus = "approved"
	DocumentStatusRejected DocumentStatus = "rejected"
)

//DocumentStatusManagerConfirmed   Status = "manager_confirmed"
//DocumentStatusEconomistConfirmed Status = "economist_confirmed"
//DocumentStatusBossConfirmed      Status = "boss_confirmed"
//DocumentStatusManagerFailed      Status = "manager_failed"
//DocumentStatusEconomistFailed    Status = "economist_failed"
//DocumentStatusBossFailed         Status = "boss_failed"

type Document struct {
	ID     int64
	Topic  string
	Status DocumentStatus
}
