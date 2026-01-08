package fhir430

import "encoding/json"

// MedicinalProductAuthorization is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductAuthorization
// The regulatory authorization of a medicinal product.
type MedicinalProductAuthorization struct {
	ID                          *string                                                    `json:"id,omitempty"`
	Meta                        *Meta                                                      `json:"meta,omitempty"`
	ImplicitRules               *string                                                    `json:"implicitRules,omitempty"`
	Language                    *string                                                    `json:"language,omitempty"`
	Text                        *Narrative                                                 `json:"text,omitempty"`
	Contained                   []json.RawMessage                                          `json:"contained,omitempty"`
	Extension                   []Extension                                                `json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `json:"modifierExtension,omitempty"`
	Identifier                  []Identifier                                               `json:"identifier,omitempty"`
	Subject                     *Reference                                                 `json:"subject,omitempty"`
	Country                     []CodeableConcept                                          `json:"country,omitempty"`
	Jurisdiction                []CodeableConcept                                          `json:"jurisdiction,omitempty"`
	Status                      *CodeableConcept                                           `json:"status,omitempty"`
	StatusDate                  *string                                                    `json:"statusDate,omitempty"`
	RestoreDate                 *string                                                    `json:"restoreDate,omitempty"`
	ValidityPeriod              *Period                                                    `json:"validityPeriod,omitempty"`
	DataExclusivityPeriod       *Period                                                    `json:"dataExclusivityPeriod,omitempty"`
	DateOfFirstAuthorization    *string                                                    `json:"dateOfFirstAuthorization,omitempty"`
	InternationalBirthDate      *string                                                    `json:"internationalBirthDate,omitempty"`
	LegalBasis                  *CodeableConcept                                           `json:"legalBasis,omitempty"`
	JurisdictionalAuthorization []MedicinalProductAuthorizationJurisdictionalAuthorization `json:"jurisdictionalAuthorization,omitempty"`
	Holder                      *Reference                                                 `json:"holder,omitempty"`
	Regulator                   *Reference                                                 `json:"regulator,omitempty"`
	Procedure                   *MedicinalProductAuthorizationProcedure                    `json:"procedure,omitempty"`
}

// Authorization in areas within a country.
type MedicinalProductAuthorizationJurisdictionalAuthorization struct {
	ID                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Identifier          []Identifier      `json:"identifier,omitempty"`
	Country             *CodeableConcept  `json:"country,omitempty"`
	Jurisdiction        []CodeableConcept `json:"jurisdiction,omitempty"`
	LegalStatusOfSupply *CodeableConcept  `json:"legalStatusOfSupply,omitempty"`
	ValidityPeriod      *Period           `json:"validityPeriod,omitempty"`
}

// The regulatory procedure for granting or amending a marketing authorization.
type MedicinalProductAuthorizationProcedure struct {
	ID                *string                                  `json:"id,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                              `json:"identifier,omitempty"`
	Type              CodeableConcept                          `json:"type"`
	DatePeriod        *Period                                  `json:"datePeriod,omitempty"`
	DateDateTime      *string                                  `json:"dateDateTime,omitempty"`
	Application       []MedicinalProductAuthorizationProcedure `json:"application,omitempty"`
}

// This function returns resource reference information
func (r MedicinalProductAuthorization) ResourceRef() (string, *string) {
	return "MedicinalProductAuthorization", r.ID
}

// This function returns resource reference information
func (r MedicinalProductAuthorization) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductAuthorization MedicinalProductAuthorization

// MarshalJSON marshals the given MedicinalProductAuthorization as JSON into a byte slice
func (r MedicinalProductAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductAuthorization
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductAuthorization: OtherMedicinalProductAuthorization(r),
		ResourceType:                       "MedicinalProductAuthorization",
	})
}

// UnmarshalMedicinalProductAuthorization unmarshals a MedicinalProductAuthorization.
func UnmarshalMedicinalProductAuthorization(b []byte) (MedicinalProductAuthorization, error) {
	var medicinalProductAuthorization MedicinalProductAuthorization
	if err := json.Unmarshal(b, &medicinalProductAuthorization); err != nil {
		return medicinalProductAuthorization, err
	}
	return medicinalProductAuthorization, nil
}
