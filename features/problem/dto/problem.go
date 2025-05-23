package problemdto

type InsertProblemRequest struct {
	Content string `json:"content" validate:"required"`
}
