
package fhir430

import "encoding/json"
// HealthcareService is documented here http://hl7.org/fhir/StructureDefinition/HealthcareService
// The details of a healthcare service available at a location.
type HealthcareService struct {
	ID                     *string                          `json:"ID,omitempty"`
	Meta                   *Meta                            `json:"meta,omitempty"`
	ImplicitRules          *string                          `json:"implicitRules,omitempty"`
	Language               *string                          `json:"language,omitempty"`
	Text                   *Narrative                       `json:"text,omitempty"`
	Contained              []json.RawMessage                `json:"contained,omitempty"`
	Extension              []Extension                      `json:"extension,omitempty"`
	ModifierExtension      []Extension                      `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                     `json:"identifier,omitempty"`
	Active                 *bool                            `json:"active,omitempty"`
	ProvidedBy             *Reference                       `json:"providedBy,omitempty"`
	Category               []CodeableConcept                `json:"category,omitempty"`
	Type                   []CodeableConcept                `json:"type,omitempty"`
	Specialty              []CodeableConcept                `json:"specialty,omitempty"`
	Location               []Reference                      `json:"location,omitempty"`
	Name                   *string                          `json:"name,omitempty"`
	Comment                *string                          `json:"comment,omitempty"`
	ExtraDetails           *string                          `json:"extraDetails,omitempty"`
	Photo                  *Attachment                      `json:"photo,omitempty"`
	Telecom                []ContactPoint                   `json:"telecom,omitempty"`
	CoverageArea           []Reference                      `json:"coverageArea,omitempty"`
	ServiceProvisionCode   []CodeableConcept                `json:"serviceProvisionCode,omitempty"`
	Eligibility            []HealthcareServiceEligibility   `json:"eligibility,omitempty"`
	Program                []CodeableConcept                `json:"program,omitempty"`
	Characteristic         []CodeableConcept                `json:"characteristic,omitempty"`
	Communication          []CodeableConcept                `json:"communication,omitempty"`
	ReferralMethod         []CodeableConcept                `json:"referralMethod,omitempty"`
	AppointmentRequired    *bool                            `json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTime `json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                          `json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                      `json:"endpoint,omitempty"`
}

// Does this service have specific eligibility requirements that need to be met in order to use the service?
type HealthcareServiceEligibility struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
}

// A collection of times that the Service Site is available.
// More detailed availability information may be provided in associated Schedule/Slot resources.
type HealthcareServiceAvailableTime struct {
	ID                 *string      `json:"ID,omitempty"`
	Extension          []Extension  `json:"extension,omitempty"`
	ModifierExtension  []Extension  `json:"modifierExtension,omitempty"`
	DaysOfWeek         []DaysOfWeek `json:"daysOfWeek,omitempty"`
	AllDay             *bool        `json:"allDay,omitempty"`
	AvailableStartTime *string      `json:"availableStartTime,omitempty"`
	AvailableEndTime   *string      `json:"availableEndTime,omitempty"`
}

// The HealthcareService is not available during this period of time due to the provided reason.
type HealthcareServiceNotAvailable struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	During            *Period     `json:"during,omitempty"`
}

// This function returns resource reference information
func (r HealthcareService) ResourceRef() (string, *string) {
	return "HealthcareService", r.ID
}

// This function returns resource reference information
func (r HealthcareService) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherHealthcareService HealthcareService

// MarshalJSON marshals the given HealthcareService as JSON into a byte slice
func (r HealthcareService) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherHealthcareService
		ResourceType string `json:"resourceType"`
	}{
		OtherHealthcareService: OtherHealthcareService(r),
		ResourceType:           "HealthcareService",
	})
}

// UnmarshalHealthcareService unmarshals a HealthcareService.
func UnmarshalHealthcareService(b []byte) (HealthcareService, error) {
	var healthcareService HealthcareService
	if err := json.Unmarshal(b, &healthcareService); err != nil {
		return healthcareService, err
	}
	return healthcareService, nil
}
