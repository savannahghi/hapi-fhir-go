package fhir500

import (
	"github.com/savannahghi/scalarutils"
)

// Duration definition: a length of time.
type Duration struct {
	ID         *string                 `json:"id,omitempty"`
	Extension  []Extension             `json:"extension,omitempty"`
	Value      *scalarutils.Decimal    `json:"value,omitempty"`
	Comparator *DurationComparatorEnum `json:"comparator,omitempty"`
	Unit       *string                 `json:"unit,omitempty"`
	System     *string                 `json:"system,omitempty"`
	Code       *string                 `json:"code,omitempty"`
}
