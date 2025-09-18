
package fhir430

import "encoding/json"
// OrganizationAffiliation is documented here http://hl7.org/fhir/StructureDefinition/OrganizationAffiliation
// Defines an affiliation/assotiation/relationship between 2 distinct oganizations, that is not a part-of relationship/sub-division relationship.
type OrganizationAffiliation struct {
	ID                        *string           `json:"ID,omitempty"`
	Meta                      *Meta             `json:"meta,omitempty"`
	ImplicitRules             *string           `json:"implicitRules,omitempty"`
	Language                  *string           `json:"language,omitempty"`
	Text                      *Narrative        `json:"text,omitempty"`
	Contained                 []json.RawMessage `json:"contained,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Identifier                []Identifier      `json:"identifier,omitempty"`
	Active                    *bool             `json:"active,omitempty"`
	Period                    *Period           `json:"period,omitempty"`
	Organization              *Reference        `json:"organization,omitempty"`
	ParticipatingOrganization *Reference        `json:"participatingOrganization,omitempty"`
	Network                   []Reference       `json:"network,omitempty"`
	Code                      []CodeableConcept `json:"code,omitempty"`
	Specialty                 []CodeableConcept `json:"specialty,omitempty"`
	Location                  []Reference       `json:"location,omitempty"`
	HealthcareService         []Reference       `json:"healthcareService,omitempty"`
	Telecom                   []ContactPoint    `json:"telecom,omitempty"`
	Endpoint                  []Reference       `json:"endpoint,omitempty"`
}

// This function returns resource reference information
func (r OrganizationAffiliation) ResourceRef() (string, *string) {
	return "OrganizationAffiliation", r.ID
}

// This function returns resource reference information
func (r OrganizationAffiliation) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherOrganizationAffiliation OrganizationAffiliation

// MarshalJSON marshals the given OrganizationAffiliation as JSON into a byte slice
func (r OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganizationAffiliation
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganizationAffiliation: OtherOrganizationAffiliation(r),
		ResourceType:                 "OrganizationAffiliation",
	})
}

// UnmarshalOrganizationAffiliation unmarshals a OrganizationAffiliation.
func UnmarshalOrganizationAffiliation(b []byte) (OrganizationAffiliation, error) {
	var organizationAffiliation OrganizationAffiliation
	if err := json.Unmarshal(b, &organizationAffiliation); err != nil {
		return organizationAffiliation, err
	}
	return organizationAffiliation, nil
}
