package fhir500

import (
	"encoding/json"

	"github.com/savannahghi/scalarutils"
)

// Immunization is documented here https://hl7.org/fhir/R5/immunization.html
// Describes the event of a patient being administered a vaccine or a record of an
// immunization as reported by a patient, a clinician or another party.
type Immunization struct {
	ID                    *string                           `json:"id,omitempty"`
	Meta                  *Meta                             `json:"meta,omitempty"`
	ImplicitRules         *string                           `json:"implicitRules,omitempty"`
	Language              *string                           `json:"language,omitempty"`
	Text                  *Narrative                        `json:"text,omitempty"`
	Extension             []Extension                       `json:"extension,omitempty"`
	ModifierExtension     []Extension                       `json:"modifierExtension,omitempty"`
	Identifier            []*Identifier                     `json:"identifier,omitempty"`
	BasedOn               []*Reference                      `json:"basedOn,omitempty"`
	Status                *ImmunizationStatus               `json:"status,omitempty"`
	StatusReason          *CodeableConcept                  `json:"statusReason,omitempty"`
	VaccineCode           *CodeableConcept                  `json:"vaccineCode,omitempty"`
	AdministeredProduct   *CodeableReference                `json:"administeredProduct,omitempty"`
	Manufacturer          *CodeableReference                `json:"manufacturer,omitempty"`
	LotNumber             *string                           `json:"lotNumber,omitempty"`
	ExpirationDate        *scalarutils.Date                 `json:"expirationDate,omitempty"`
	Patient               *Reference                        `json:"patient,omitempty"`
	Encounter             *Reference                        `json:"encounter,omitempty"`
	SupportingInformation []*Reference                      `json:"supportingInformation,omitempty"`
	OccurrenceDateTime    *scalarutils.Date                 `json:"occurrenceDateTime,omitempty"`
	OccurrenceString      *string                           `json:"occurrenceString,omitempty"`
	PrimarySource         *bool                             `json:"primarySource,omitempty"`
	InformationSource     *CodeableReference                `json:"informationSource,omitempty"`
	Location              *Reference                        `json:"location,omitempty"`
	Site                  *CodeableConcept                  `json:"site,omitempty"`
	Route                 *CodeableConcept                  `json:"route,omitempty"`
	DoseQuantity          *Quantity                         `json:"doseQuantity,omitempty"`
	Performer             []*ImmunizationPerformer          `json:"performer,omitempty"`
	Note                  []*Annotation                     `json:"note,omitempty"`
	Reason                []*CodeableReference              `json:"reason,omitempty"`
	IsSubpotent           *bool                             `json:"isSubpotent,omitempty"`
	SubpotentReason       []*CodeableConcept                `json:"subpotentReason,omitempty"`
	ProgramEligibility    []*ImmunizationProgramEligibility `json:"programEligibility,omitempty"`
	FundingSource         *CodeableConcept                  `json:"fundingSource,omitempty"`
	Reaction              []*ImmunizationReaction           `json:"reaction,omitempty"`
	ProtocolApplied       []*ImmunizationProtocolApplied    `json:"protocolApplied,omitempty"`
}

// ImmunizationPerformer indicates who performed the immunization event.
type ImmunizationPerformer struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             *Reference       `json:"actor,omitempty"`
}

// ImmunizationProgramEligibility describes the patient's eligibility for a
// vaccination program.
//
// R5 promoted programEligibility from a plain CodeableConcept to a backbone
// element so the program and the patient's status within it are recorded
// separately.
type ImmunizationProgramEligibility struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Program           *CodeableConcept `json:"program,omitempty"`
	ProgramStatus     *CodeableConcept `json:"programStatus,omitempty"`
}

// ImmunizationReaction holds categorical data indicating that an adverse event
// is associated in time to an immunization.
//
// A reaction may indicate an allergy or intolerance. Where that is established,
// record it as a new AllergyIntolerance instance, because most systems do not
// query past Immunization.reaction elements.
type ImmunizationReaction struct {
	ID                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Date              *scalarutils.Date  `json:"date,omitempty"`
	Manifestation     *CodeableReference `json:"manifestation,omitempty"`
	Reported          *bool              `json:"reported,omitempty"`
}

// ImmunizationProtocolApplied is the protocol, meaning the set of
// recommendations, followed by the provider who administered the dose.
type ImmunizationProtocolApplied struct {
	ID                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Series            *string            `json:"series,omitempty"`
	Authority         *Reference         `json:"authority,omitempty"`
	TargetDisease     []*CodeableConcept `json:"targetDisease,omitempty"`
	DoseNumber        *string            `json:"doseNumber,omitempty"`
	SeriesDoses       *string            `json:"seriesDoses,omitempty"`
}

// ImmunizationRelayPayload is the wrapper used when an Immunization is relayed
// inside another payload.
type ImmunizationRelayPayload struct {
	Resource *Immunization `json:"resource,omitempty"`
}

type OtherImmunization Immunization

// MarshalJSON marshals the given Immunization as JSON into a byte slice.
func (r Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunization
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunization: OtherImmunization(r),
		ResourceType:      "Immunization",
	})
}

// UnmarshalImmunization unmarshals an Immunization.
func UnmarshalImmunization(b []byte) (Immunization, error) {
	var immunization Immunization
	if err := json.Unmarshal(b, &immunization); err != nil {
		return immunization, err
	}

	return immunization, nil
}
