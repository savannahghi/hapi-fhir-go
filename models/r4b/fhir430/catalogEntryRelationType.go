
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// CatalogEntryRelationType is documented here http://hl7.org/fhir/ValueSet/relation-type
type CatalogEntryRelationType int

const (
	CatalogEntryRelationTypeTriggers CatalogEntryRelationType = iota
	CatalogEntryRelationTypeIsReplacedBy
)

func (code CatalogEntryRelationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CatalogEntryRelationType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal CatalogEntryRelationType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "triggers":
		*code = CatalogEntryRelationTypeTriggers
	case "is-replaced-by":
		*code = CatalogEntryRelationTypeIsReplacedBy
	default:
		return fmt.Errorf("unknown CatalogEntryRelationType code `%s`", s)
	}
	return nil
}
func (code CatalogEntryRelationType) String() string {
	return code.Code()
}
func (code CatalogEntryRelationType) Code() string {
	switch code {
	case CatalogEntryRelationTypeTriggers:
		return "triggers"
	case CatalogEntryRelationTypeIsReplacedBy:
		return "is-replaced-by"
	}
	return "<unknown>"
}
func (code CatalogEntryRelationType) Display() string {
	switch code {
	case CatalogEntryRelationTypeTriggers:
		return "Triggers"
	case CatalogEntryRelationTypeIsReplacedBy:
		return "Replaced By"
	}
	return "<unknown>"
}
func (code CatalogEntryRelationType) Definition() string {
	switch code {
	case CatalogEntryRelationTypeTriggers:
		return "the related entry represents an activity that may be triggered by the current item."
	case CatalogEntryRelationTypeIsReplacedBy:
		return "the related entry represents an item that replaces the current retired item."
	}
	return "<unknown>"
}
