package sqlite_test

import (
	"context"
	"testing"

	"github.com/stillwondering/themis"
	"github.com/stillwondering/themis/sqlite"
)

func TestCreateEvent(t *testing.T) {
	db := MustOpenDB(t)
	defer MustCloseDB(t, db)
	s := sqlite.NewEventService(db)
	s.GenerateUUID = func() string {
		return "abcd"
	}

	event, err := s.CreateEvent(context.Background(), themis.EventCreate{
		Title:       "First event",
		Description: "Just some description",
	})

	if err != nil {
		t.Fatal(err)
	}

	expectedEvent := themis.Event{
		ID:          1,
		UUID:        "abcd",
		Title:       "First event",
		Description: "Just some description",
	}

	if expectedEvent != *event {
		t.Fatalf("expected: %v\ngot: %v", expectedEvent, *event)
	}
}
