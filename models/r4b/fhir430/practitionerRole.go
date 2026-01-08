package fhir430

import "encoding/json"

// PractitionerRole is documented here http://hl7.org/fhir/StructureDefinition/PractitionerRole
// A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time.
type PractitionerRole struct {
	ID                     *string                         `json:"id,omitempty"`
	Meta                   *Meta                           `json:"meta,omitempty"`
	ImplicitRules          *string                         `json:"implicitRules,omitempty"`
	Language               *string                         `json:"language,omitempty"`
	Text                   *Narrative                      `json:"text,omitempty"`
	Contained              []json.RawMessage               `json:"contained,omitempty"`
	Extension              []Extension                     `json:"extension,omitempty"`
	ModifierExtension      []Extension                     `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                    `json:"identifier,omitempty"`
	Active                 *bool                           `json:"active,omitempty"`
	Period                 *Period                         `json:"period,omitempty"`
	Practitioner           *Reference                      `json:"practitioner,omitempty"`
	Organization           *Reference                      `json:"organization,omitempty"`
	Code                   []CodeableConcept               `json:"code,omitempty"`
	Specialty              []CodeableConcept               `json:"specialty,omitempty"`
	Location               []Reference                     `json:"location,omitempty"`
	HealthcareService      []Reference                     `json:"healthcareService,omitempty"`
	Telecom                []ContactPoint                  `json:"telecom,omitempty"`
	AvailableTime          []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`
	NotAvailable           []PractitionerRoleNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                         `json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                     `json:"endpoint,omitempty"`
}

// A collection of times the practitioner is available or performing this role at the location and/or healthcareservice.
// More detailed availability information may be provided in associated Schedule/Slot resources.
type PractitionerRoleAvailableTime struct {
	ID                 *string      `json:"id,omitempty"`
	Extension          []Extension  `json:"extension,omitempty"`
	ModifierExtension  []Extension  `json:"modifierExtension,omitempty"`
	DaysOfWeek         []DaysOfWeek `json:"daysOfWeek,omitempty"`
	AllDay             *bool        `json:"allDay,omitempty"`
	AvailableStartTime *string      `json:"availableStartTime,omitempty"`
	AvailableEndTime   *string      `json:"availableEndTime,omitempty"`
}

// The practitioner is not available or performing this role during this period of time due to the provided reason.
type PractitionerRoleNotAvailable struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	During            *Period     `json:"during,omitempty"`
}

// This function returns resource reference information
func (r PractitionerRole) ResourceRef() (string, *string) {
	return "PractitionerRole", r.ID
}

// This function returns resource reference information
func (r PractitionerRole) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherPractitionerRole PractitionerRole

// MarshalJSON marshals the given PractitionerRole as JSON into a byte slice
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}

// UnmarshalPractitionerRole unmarshals a PractitionerRole.
func UnmarshalPractitionerRole(b []byte) (PractitionerRole, error) {
	var practitionerRole PractitionerRole
	if err := json.Unmarshal(b, &practitionerRole); err != nil {
		return practitionerRole, err
	}
	return practitionerRole, nil
}
