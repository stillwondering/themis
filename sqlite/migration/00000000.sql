CREATE TABLE events (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid        TEXT NOT NULL UNIQUE,
    title       TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE enrollments (
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid     TEXT NOT NULL UNIQUE,
    event_id INTEGER NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    comment  TEXT DEFAULT ""
);

CREATE TABLE enrolled_people (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    enrollment_id INTEGER NOT NULL REFERENCES enrollments (id) ON DELETE CASCADE,
    name          TEXT NOT NULL
);

CREATE INDEX events_uuid_idx ON events (uuid);
CREATE INDEX enrollments_uuid_idx ON enrollments (uuid);
CREATE INDEX enrollments_event_id_idx ON enrollments (event_id);
CREATE INDEX enrolled_people_enrollment_id_idx ON enrolled_people (enrollment_id);