package campaigns

const (
	StatusActive = "active"
	StatusPaused = "paused"
)

type Campaign struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`

	StartTs int64 `json:"start_ts"`
	EndTs   int64 `json:"end_ts"`

	VisitsGoal int32 `json:"visits_goal"`
}
