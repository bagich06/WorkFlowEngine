package entities

type DocumentStatus string

const (
	DocumentStatusStarted            DocumentStatus = "started"
	DocumentStatusManagerConfirmed   DocumentStatus = "manager_confirmed"
	DocumentStatusEconomistConfirmed DocumentStatus = "economist_confirmed"
	DocumentStatusBossConfirmed      DocumentStatus = "boss_confirmed"
	DocumentStatusManagerFailed      DocumentStatus = "manager_failed"
	DocumentStatusEconomistFailed    DocumentStatus = "economist_failed"
	DocumentStatusBossFailed         DocumentStatus = "boss_failed"
	DocumentStatusExpired            DocumentStatus = "expired"
	DocumentStatusSuccess            DocumentStatus = "success"
)

type Document struct {
	ID     int64
	Topic  string
	Status DocumentStatus
}
