package form

type Lottery struct {
	URL         string `json:"url" validate:"required"`
	WinnerCount int    `json:"winnerCount" validate:"required,gt=0"`
}
