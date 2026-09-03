package consent

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		in   string
		want time.Time
	}{
		{"2026", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"2026-09", time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)},
		{"2026-09-03", time.Date(2026, 9, 3, 0, 0, 0, 0, time.UTC)},
		{"2026-09-03T08:00:00+03:00", time.Date(2026, 9, 3, 5, 0, 0, 0, time.UTC)},
		{"2026-09-03T08:00:00.123456+03:00", time.Date(2026, 9, 3, 5, 0, 0, 123456000, time.UTC)},
		{"2026-09-03T08:00:00Z", time.Date(2026, 9, 3, 8, 0, 0, 0, time.UTC)},
		{"2026-09-03T08:00:00", time.Date(2026, 9, 3, 8, 0, 0, 0, time.UTC)},
	}

	for _, tc := range tests {
		got, err := ParseTime(tc.in)
		if err != nil {
			t.Errorf("ParseTime(%q) error: %v", tc.in, err)

			continue
		}

		if !got.Equal(tc.want) {
			t.Errorf("ParseTime(%q) = %s, want %s", tc.in, got, tc.want)
		}
	}

	if _, err := ParseTime("yesterday"); err == nil {
		t.Error("ParseTime(\"yesterday\") expected error")
	}
}

func TestParsePeriodEnd(t *testing.T) {
	tests := []struct {
		in   string
		want time.Time
	}{
		{"2026", time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)},
		{"2026-02", time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)},
		{"2026-01-31", time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)},
		{"2026-01-31T10:00:00Z", time.Date(2026, 1, 31, 10, 0, 0, 0, time.UTC)},
	}

	for _, tc := range tests {
		got, err := ParsePeriodEnd(tc.in)
		if err != nil {
			t.Errorf("ParsePeriodEnd(%q) error: %v", tc.in, err)

			continue
		}

		if !got.Equal(tc.want) {
			t.Errorf("ParsePeriodEnd(%q) = %s, want %s", tc.in, got, tc.want)
		}
	}
}
