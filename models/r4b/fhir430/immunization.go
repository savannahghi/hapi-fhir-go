package fhir430

import "encoding/json"

// Immunization is documented here http://hl7.org/fhir/StructureDefinition/Immunization
// Describes the event of a patient being administered a vaccine or a record of an immunization as reported by a patient, a clinician or another party.
type Immunization struct {
	ID                 *string                       `json:"id,omitempty"`
	Meta               *Meta                         `json:"meta,omitempty"`
	ImplicitRules      *string                       `json:"implicitRules,omitempty"`
	Language           *string                       `json:"language,omitempty"`
	Text               *Narrative                    `json:"text,omitempty"`
	Contained          []json.RawMessage             `json:"contained,omitempty"`
	Extension          []Extension                   `json:"extension,omitempty"`
	ModifierExtension  []Extension                   `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                  `json:"identifier,omitempty"`
	Status             ImmunizationStatusCodes       `json:"status"`
	StatusReason       *CodeableConcept              `json:"statusReason,omitempty"`
	VaccineCode        CodeableConcept               `json:"vaccineCode"`
	Patient            Reference                     `json:"patient"`
	Encounter          *Reference                    `json:"encounter,omitempty"`
	OccurrenceDateTime string                        `json:"occurrenceDateTime"`
	OccurrenceString   string                        `json:"occurrenceString"`
	Recorded           *string                       `json:"recorded,omitempty"`
	PrimarySource      *bool                         `json:"primarySource,omitempty"`
	ReportOrigin       *CodeableConcept              `json:"reportOrigin,omitempty"`
	Location           *Reference                    `json:"location,omitempty"`
	Manufacturer       *Reference                    `json:"manufacturer,omitempty"`
	LotNumber          *string                       `json:"lotNumber,omitempty"`
	ExpirationDate     *string                       `json:"expirationDate,omitempty"`
	Site               *CodeableConcept              `json:"site,omitempty"`
	Route              *CodeableConcept              `json:"route,omitempty"`
	DoseQuantity       *Quantity                     `json:"doseQuantity,omitempty"`
	Performer          []ImmunizationPerformer       `json:"performer,omitempty"`
	Note               []Annotation                  `json:"note,omitempty"`
	ReasonCode         []CodeableConcept             `json:"reasonCode,omitempty"`
	ReasonReference    []Reference                   `json:"reasonReference,omitempty"`
	IsSubpotent        *bool                         `json:"isSubpotent,omitempty"`
	SubpotentReason    []CodeableConcept             `json:"subpotentReason,omitempty"`
	Education          []ImmunizationEducation       `json:"education,omitempty"`
	ProgramEligibility []CodeableConcept             `json:"programEligibility,omitempty"`
	FundingSource      *CodeableConcept              `json:"fundingSource,omitempty"`
	Reaction           []ImmunizationReaction        `json:"reaction,omitempty"`
	ProtocolApplied    []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`
}

// Indicates who performed the immunization event.
type ImmunizationPerformer struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// Educational material presented to the patient (or guardian) at the time of vaccine administration.
type ImmunizationEducation struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DocumentType      *string     `json:"documentType,omitempty"`
	Reference         *string     `json:"reference,omitempty"`
	PublicationDate   *string     `json:"publicationDate,omitempty"`
	PresentationDate  *string     `json:"presentationDate,omitempty"`
}

// Categorical data indicating that an adverse event is associated in time to an immunization.
// A reaction may be an indication of an allergy or intolerance and, if this is determined to be the case, it should be recorded as a new AllergyIntolerance resource instance as most systems will not query against past Immunization.reaction elements.
type ImmunizationReaction struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Date              *string     `json:"date,omitempty"`
	Detail            *Reference  `json:"detail,omitempty"`
	Reported          *bool       `json:"reported,omitempty"`
}

// The protocol (set of recommendations) being followed by the provider who administered the dose.
type ImmunizationProtocolApplied struct {
	ID                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Series                 *string           `json:"series,omitempty"`
	Authority              *Reference        `json:"authority,omitempty"`
	TargetDisease          []CodeableConcept `json:"targetDisease,omitempty"`
	DoseNumberPositiveInt  int               `json:"doseNumberPositiveInt"`
	DoseNumberString       string            `json:"doseNumberString"`
	SeriesDosesPositiveInt *int              `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString      *string           `json:"seriesDosesString,omitempty"`
}

// This function returns resource reference information
func (r Immunization) ResourceRef() (string, *string) {
	return "Immunization", r.ID
}

// This function returns resource reference information
func (r Immunization) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherImmunization Immunization

// MarshalJSON marshals the given Immunization as JSON into a byte slice
func (r Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunization
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunization: OtherImmunization(r),
		ResourceType:      "Immunization",
	})
}

// UnmarshalImmunization unmarshals a Immunization.
func UnmarshalImmunization(b []byte) (Immunization, error) {
	var immunization Immunization
	if err := json.Unmarshal(b, &immunization); err != nil {
		return immunization, err
	}
	return immunization, nil
}
