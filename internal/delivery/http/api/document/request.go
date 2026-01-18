package document

type DocumentCreateRequest struct {
	Topic string `json:"topic" validate:"required"`
}
