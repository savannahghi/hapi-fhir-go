package consent

import (
	"fmt"
	"time"
)

// FHIR date and dateTime layouts, most precise first.
var fhirLayouts = []struct {
	layout string
	// unit is the calendar unit a value at this precision covers; used to
	// widen an inclusive period end.
	unit func(time.Time) time.Time
}{
	{time.RFC3339Nano, nil},
	{time.RFC3339, nil},
	{"2006-01-02T15:04:05.999999999", nil},
	{"2006-01-02T15:04:05", nil},
	{"2006-01-02", func(t time.Time) time.Time { return t.AddDate(0, 0, 1) }},
	{"2006-01", func(t time.Time) time.Time { return t.AddDate(0, 1, 0) }},
	{"2006", func(t time.Time) time.Time { return t.AddDate(1, 0, 0) }},
}

// ParseTime parses a FHIR date, dateTime or instant. Partial dates resolve to
// the start of the period they name, in UTC. Times without an offset are read
// as UTC.
func ParseTime(s string) (time.Time, error) {
	t, _, err := parseFHIR(s)

	return t, err
}

// ParsePeriodEnd parses the inclusive end of a FHIR Period. A partial date
// covers the whole unit it names, so "2026-01-31" ends just before
// 2026-02-01T00:00:00Z.
func ParsePeriodEnd(s string) (time.Time, error) {
	t, widen, err := parseFHIR(s)
	if err != nil {
		return t, err
	}

	if widen != nil {
		return widen(t).Add(-time.Nanosecond), nil
	}

	return t, nil
}

func parseFHIR(s string) (time.Time, func(time.Time) time.Time, error) {
	for _, l := range fhirLayouts {
		t, err := time.ParseInLocation(l.layout, s, time.UTC)
		if err == nil {
			return t, l.unit, nil
		}
	}

	return time.Time{}, nil, fmt.Errorf("invalid FHIR date/dateTime %q", s)
}
