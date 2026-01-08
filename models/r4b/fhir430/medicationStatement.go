package fhir430

import "encoding/json"

// MedicationStatement is documented here http://hl7.org/fhir/StructureDefinition/MedicationStatement
/*
A record of a medication that is being consumed by a patient.   A MedicationStatement may indicate that the patient may be taking the medication now or has taken the medication in the past or will be taking the medication in the future.  The source of this information can be the patient, significant other (such as a family member or spouse), or a clinician.  A common scenario where this information is captured is during the history taking process during a patient visit or stay.   The medication information may come from sources such as the patient's memory, from a prescription bottle,  or from a list of medications the patient, clinician or other party maintains.

The primary difference between a medication statement and a medication administration is that the medication administration has complete administration information and is based on actual administration information from the person who administered the medication.  A medication statement is often, if not always, less specific.  There is no required date/time when the medication was administered, in fact we only know that a source has reported the patient is taking this medication, where details such as time, quantity, or rate or even medication product may be incomplete or missing or less precise.  As stated earlier, the medication statement information may come from the patient's memory, from a prescription bottle or from a list of medications the patient, clinician or other party maintains.  Medication administration is more formal and is not missing detailed information.
*/
type MedicationStatement struct {
	ID                        *string           `json:"id,omitempty"`
	Meta                      *Meta             `json:"meta,omitempty"`
	ImplicitRules             *string           `json:"implicitRules,omitempty"`
	Language                  *string           `json:"language,omitempty"`
	Text                      *Narrative        `json:"text,omitempty"`
	Contained                 []json.RawMessage `json:"contained,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Identifier                []Identifier      `json:"identifier,omitempty"`
	BasedOn                   []Reference       `json:"basedOn,omitempty"`
	PartOf                    []Reference       `json:"partOf,omitempty"`
	Status                    string            `json:"status"`
	StatusReason              []CodeableConcept `json:"statusReason,omitempty"`
	Category                  *CodeableConcept  `json:"category,omitempty"`
	MedicationCodeableConcept CodeableConcept   `json:"medicationCodeableConcept"`
	MedicationReference       Reference         `json:"medicationReference"`
	Subject                   Reference         `json:"subject"`
	Context                   *Reference        `json:"context,omitempty"`
	EffectiveDateTime         *string           `json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period           `json:"effectivePeriod,omitempty"`
	DateAsserted              *string           `json:"dateAsserted,omitempty"`
	InformationSource         *Reference        `json:"informationSource,omitempty"`
	DerivedFrom               []Reference       `json:"derivedFrom,omitempty"`
	ReasonCode                []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference           []Reference       `json:"reasonReference,omitempty"`
	Note                      []Annotation      `json:"note,omitempty"`
	Dosage                    []Dosage          `json:"dosage,omitempty"`
}

// This function returns resource reference information
func (r MedicationStatement) ResourceRef() (string, *string) {
	return "MedicationStatement", r.ID
}

// This function returns resource reference information
func (r MedicationStatement) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicationStatement MedicationStatement

// MarshalJSON marshals the given MedicationStatement as JSON into a byte slice
func (r MedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationStatement: OtherMedicationStatement(r),
		ResourceType:             "MedicationStatement",
	})
}

// UnmarshalMedicationStatement unmarshals a MedicationStatement.
func UnmarshalMedicationStatement(b []byte) (MedicationStatement, error) {
	var medicationStatement MedicationStatement
	if err := json.Unmarshal(b, &medicationStatement); err != nil {
		return medicationStatement, err
	}
	return medicationStatement, nil
}
