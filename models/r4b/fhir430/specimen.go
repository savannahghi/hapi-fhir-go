
package fhir430

import "encoding/json"
// Specimen is documented here http://hl7.org/fhir/StructureDefinition/Specimen
// A sample to be used for analysis.
type Specimen struct {
	ID                  *string              `json:"ID,omitempty"`
	Meta                *Meta                `json:"meta,omitempty"`
	ImplicitRules       *string              `json:"implicitRules,omitempty"`
	Language            *string              `json:"language,omitempty"`
	Text                *Narrative           `json:"text,omitempty"`
	Contained           []json.RawMessage    `json:"contained,omitempty"`
	Extension           []Extension          `json:"extension,omitempty"`
	ModifierExtension   []Extension          `json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `json:"identifier,omitempty"`
	AccessionIdentifier *Identifier          `json:"accessionIdentifier,omitempty"`
	Status              *SpecimenStatus      `json:"status,omitempty"`
	Type                *CodeableConcept     `json:"type,omitempty"`
	Subject             *Reference           `json:"subject,omitempty"`
	ReceivedTime        *string              `json:"receivedTime,omitempty"`
	Parent              []Reference          `json:"parent,omitempty"`
	Request             []Reference          `json:"request,omitempty"`
	Collection          *SpecimenCollection  `json:"collection,omitempty"`
	Processing          []SpecimenProcessing `json:"processing,omitempty"`
	Container           []SpecimenContainer  `json:"container,omitempty"`
	Condition           []CodeableConcept    `json:"condition,omitempty"`
	Note                []Annotation         `json:"note,omitempty"`
}

// Details concerning the specimen collection.
type SpecimenCollection struct {
	ID                           *string          `json:"ID,omitempty"`
	Extension                    []Extension      `json:"extension,omitempty"`
	ModifierExtension            []Extension      `json:"modifierExtension,omitempty"`
	Collector                    *Reference       `json:"collector,omitempty"`
	CollectedDateTime            *string          `json:"collectedDateTime,omitempty"`
	CollectedPeriod              *Period          `json:"collectedPeriod,omitempty"`
	Duration                     *Duration        `json:"duration,omitempty"`
	Quantity                     *Quantity        `json:"quantity,omitempty"`
	Method                       *CodeableConcept `json:"method,omitempty"`
	BodySite                     *CodeableConcept `json:"bodySite,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration        `json:"fastingStatusDuration,omitempty"`
}

// Details concerning processing and processing steps for the specimen.
type SpecimenProcessing struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          []Reference      `json:"additive,omitempty"`
	TimeDateTime      *string          `json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `json:"timePeriod,omitempty"`
}

// The container holding the specimen.  The recursive nature of containers; i.e. blood in tube in tray in rack is not addressed here.
type SpecimenContainer struct {
	ID                      *string          `json:"ID,omitempty"`
	Extension               []Extension      `json:"extension,omitempty"`
	ModifierExtension       []Extension      `json:"modifierExtension,omitempty"`
	Identifier              []Identifier     `json:"identifier,omitempty"`
	Description             *string          `json:"description,omitempty"`
	Type                    *CodeableConcept `json:"type,omitempty"`
	Capacity                *Quantity        `json:"capacity,omitempty"`
	SpecimenQuantity        *Quantity        `json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `json:"additiveReference,omitempty"`
}

// This function returns resource reference information
func (r Specimen) ResourceRef() (string, *string) {
	return "Specimen", r.ID
}

// This function returns resource reference information
func (r Specimen) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSpecimen Specimen

// MarshalJSON marshals the given Specimen as JSON into a byte slice
func (r Specimen) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimen
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimen: OtherSpecimen(r),
		ResourceType:  "Specimen",
	})
}

// UnmarshalSpecimen unmarshals a Specimen.
func UnmarshalSpecimen(b []byte) (Specimen, error) {
	var specimen Specimen
	if err := json.Unmarshal(b, &specimen); err != nil {
		return specimen, err
	}
	return specimen, nil
}
