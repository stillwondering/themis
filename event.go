package themis

// Event represents a certain event that people can attend.
type Event struct {
	ID          int
	UUID        string
	Title       string
	Description string
}

type EventCreate struct {
	Title       string
	Description string
}
