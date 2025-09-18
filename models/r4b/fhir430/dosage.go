
package fhir430
// Dosage is documented here http://hl7.org/fhir/StructureDefinition/Dosage
// Base StructureDefinition for Dosage Type: Indicates how the medication is/was taken or should be taken by the patient.
type Dosage struct {
	ID                       *string             `json:"ID,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	Sequence                 *int                `json:"sequence,omitempty"`
	Text                     *string             `json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept   `json:"additionalInstruction,omitempty"`
	PatientInstruction       *string             `json:"patientInstruction,omitempty"`
	Timing                   *Timing             `json:"timing,omitempty"`
	AsNeededBoolean          *bool               `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept  *CodeableConcept    `json:"asNeededCodeableConcept,omitempty"`
	Site                     *CodeableConcept    `json:"site,omitempty"`
	Route                    *CodeableConcept    `json:"route,omitempty"`
	Method                   *CodeableConcept    `json:"method,omitempty"`
	DoseAndRate              []DosageDoseAndRate `json:"doseAndRate,omitempty"`
	MaxDosePerPeriod         *Ratio              `json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity           `json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity           `json:"maxDosePerLifetime,omitempty"`
}

// The amount of medication administered.
type DosageDoseAndRate struct {
	ID           *string          `json:"ID,omitempty"`
	Extension    []Extension      `json:"extension,omitempty"`
	Type         *CodeableConcept `json:"type,omitempty"`
	DoseRange    *Range           `json:"doseRange,omitempty"`
	DoseQuantity *Quantity        `json:"doseQuantity,omitempty"`
	RateRatio    *Ratio           `json:"rateRatio,omitempty"`
	RateRange    *Range           `json:"rateRange,omitempty"`
	RateQuantity *Quantity        `json:"rateQuantity,omitempty"`
}
