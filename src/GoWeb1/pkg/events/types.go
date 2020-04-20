package events

type Event struct {
	ID          string `json:"id, omitempty" `
	Name        string `json:"name, omitempty"`
	Description string `json:"description, omitempty"`
}

const EventCollection = "Event"
