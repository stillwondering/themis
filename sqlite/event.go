package sqlite

import (
	"context"

	"github.com/google/uuid"
	"github.com/stillwondering/themis"
)

// EventService is a service for managing events.
type EventService struct {
	db           *DB
	GenerateUUID func() string
}

// NewEventService creates a new instance of EventService.
func NewEventService(db *DB) *EventService {
	return &EventService{
		db: db,
		GenerateUUID: func() string {
			return uuid.NewString()
		},
	}
}

func (s *EventService) FindByUUID(ctx context.Context, uuid string) (*themis.Event, error) {
	e := themis.Event{}

	row := s.db.db.QueryRowContext(ctx, `
		SELECT
			id,
			uuid,
			title,
			description
		FROM
			events
		WHERE
			uuid = ?
	`, uuid)

	if err := row.Scan(&e.ID, &e.UUID, &e.Title, &e.Description); err != nil {
		return nil, err
	}

	return &e, nil
}

func (s *EventService) CreateEvent(ctx context.Context, e themis.EventCreate) (*themis.Event, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	uuid := s.GenerateUUID()

	event, err := createEvent(ctx, tx, e, uuid)
	if err != nil {
		return nil, err
	}

	return event, tx.Commit()
}

// ListAll returns a list of all events in the database.
func (s *EventService) ListAll(ctx context.Context) ([]*themis.Event, error) {
	rows, err := s.db.db.QueryContext(ctx, `
		SELECT
			id,
			uuid,
			title,
			description
		FROM
			events
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var uuid, title, description string
	var events []*themis.Event

	for rows.Next() {
		if err := rows.Scan(&id, &uuid, &title, &description); err != nil {
			return nil, err
		}

		events = append(events, &themis.Event{
			ID:          id,
			UUID:        uuid,
			Title:       title,
			Description: description,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func createEvent(ctx context.Context, tx *Tx, data themis.EventCreate, uuid string) (*themis.Event, error) {
	result, err := tx.ExecContext(ctx, `
		INSERT INTO events (
			uuid,
			title,
			description
		) VALUES (
			?,
			?,
			?
		)
	`, uuid, data.Title, data.Description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	newEvent := themis.Event{
		ID:          int(id),
		UUID:        uuid,
		Title:       data.Title,
		Description: data.Description,
	}

	return &newEvent, nil
}
