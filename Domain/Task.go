package Domain

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"` // "todo", "in_progress", "done"
}
