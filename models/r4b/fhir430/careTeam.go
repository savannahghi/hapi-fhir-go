
package fhir430

import "encoding/json"
// CareTeam is documented here http://hl7.org/fhir/StructureDefinition/CareTeam
// The Care Team includes all the people and organizations who plan to participate in the coordination and delivery of care for a patient.
type CareTeam struct {
	ID                   *string               `json:"ID,omitempty"`
	Meta                 *Meta                 `json:"meta,omitempty"`
	ImplicitRules        *string               `json:"implicitRules,omitempty"`
	Language             *string               `json:"language,omitempty"`
	Text                 *Narrative            `json:"text,omitempty"`
	Contained            []json.RawMessage     `json:"contained,omitempty"`
	Extension            []Extension           `json:"extension,omitempty"`
	ModifierExtension    []Extension           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `json:"identifier,omitempty"`
	Status               *CareTeamStatus       `json:"status,omitempty"`
	Category             []CodeableConcept     `json:"category,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	Subject              *Reference            `json:"subject,omitempty"`
	Encounter            *Reference            `json:"encounter,omitempty"`
	Period               *Period               `json:"period,omitempty"`
	Participant          []CareTeamParticipant `json:"participant,omitempty"`
	ReasonCode           []CodeableConcept     `json:"reasonCode,omitempty"`
	ReasonReference      []Reference           `json:"reasonReference,omitempty"`
	ManagingOrganization []Reference           `json:"managingOrganization,omitempty"`
	Telecom              []ContactPoint        `json:"telecom,omitempty"`
	Note                 []Annotation          `json:"note,omitempty"`
}

// Identifies all people and organizations who are expected to be involved in the care team.
type CareTeamParticipant struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Member            *Reference        `json:"member,omitempty"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
	Period            *Period           `json:"period,omitempty"`
}

// This function returns resource reference information
func (r CareTeam) ResourceRef() (string, *string) {
	return "CareTeam", r.ID
}

// This function returns resource reference information
func (r CareTeam) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCareTeam CareTeam

// MarshalJSON marshals the given CareTeam as JSON into a byte slice
func (r CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCareTeam
		ResourceType string `json:"resourceType"`
	}{
		OtherCareTeam: OtherCareTeam(r),
		ResourceType:  "CareTeam",
	})
}

// UnmarshalCareTeam unmarshals a CareTeam.
func UnmarshalCareTeam(b []byte) (CareTeam, error) {
	var careTeam CareTeam
	if err := json.Unmarshal(b, &careTeam); err != nil {
		return careTeam, err
	}
	return careTeam, nil
}
