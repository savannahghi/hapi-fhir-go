package fhir430

// MarketingStatus is documented here http://hl7.org/fhir/StructureDefinition/MarketingStatus
// Base StructureDefinition for MarketingStatus Type: The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type MarketingStatus struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Country           CodeableConcept  `json:"country"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
	Status            CodeableConcept  `json:"status"`
	DateRange         Period           `json:"dateRange"`
	RestoreDate       *string          `json:"restoreDate,omitempty"`
}
