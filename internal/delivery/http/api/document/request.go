package document

type DocumentCreateRequest struct {
	Topic  string  `json:"topic" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}
