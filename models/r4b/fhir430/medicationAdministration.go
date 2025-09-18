
package fhir430

import "encoding/json"
// MedicationAdministration is documented here http://hl7.org/fhir/StructureDefinition/MedicationAdministration
// Describes the event of a patient consuming or otherwise being administered a medication.  This may be as simple as swallowing a tablet or it may be a long running infusion.  Related resources tie this event to the authorizing prescription, and the specific encounter between patient and health care practitioner.
type MedicationAdministration struct {
	ID                        *string                             `json:"ID,omitempty"`
	Meta                      *Meta                               `json:"meta,omitempty"`
	ImplicitRules             *string                             `json:"implicitRules,omitempty"`
	Language                  *string                             `json:"language,omitempty"`
	Text                      *Narrative                          `json:"text,omitempty"`
	Contained                 []json.RawMessage                   `json:"contained,omitempty"`
	Extension                 []Extension                         `json:"extension,omitempty"`
	ModifierExtension         []Extension                         `json:"modifierExtension,omitempty"`
	Identifier                []Identifier                        `json:"identifier,omitempty"`
	Instantiates              []string                            `json:"instantiates,omitempty"`
	PartOf                    []Reference                         `json:"partOf,omitempty"`
	Status                    string                              `json:"status"`
	StatusReason              []CodeableConcept                   `json:"statusReason,omitempty"`
	Category                  *CodeableConcept                    `json:"category,omitempty"`
	MedicationCodeableConcept CodeableConcept                     `json:"medicationCodeableConcept"`
	MedicationReference       Reference                           `json:"medicationReference"`
	Subject                   Reference                           `json:"subject"`
	Context                   *Reference                          `json:"context,omitempty"`
	SupportingInformation     []Reference                         `json:"supportingInformation,omitempty"`
	EffectiveDateTime         string                              `json:"effectiveDateTime"`
	EffectivePeriod           Period                              `json:"effectivePeriod"`
	Performer                 []MedicationAdministrationPerformer `json:"performer,omitempty"`
	ReasonCode                []CodeableConcept                   `json:"reasonCode,omitempty"`
	ReasonReference           []Reference                         `json:"reasonReference,omitempty"`
	Request                   *Reference                          `json:"request,omitempty"`
	Device                    []Reference                         `json:"device,omitempty"`
	Note                      []Annotation                        `json:"note,omitempty"`
	Dosage                    *MedicationAdministrationDosage     `json:"dosage,omitempty"`
	EventHistory              []Reference                         `json:"eventHistory,omitempty"`
}

// Indicates who or what performed the medication administration and how they were involved.
type MedicationAdministrationPerformer struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// Describes the medication dosage information details e.g. dose, rate, site, route, etc.
type MedicationAdministrationDosage struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Text              *string          `json:"text,omitempty"`
	Site              *CodeableConcept `json:"site,omitempty"`
	Route             *CodeableConcept `json:"route,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Dose              *Quantity        `json:"dose,omitempty"`
	RateRatio         *Ratio           `json:"rateRatio,omitempty"`
	RateQuantity      *Quantity        `json:"rateQuantity,omitempty"`
}

// This function returns resource reference information
func (r MedicationAdministration) ResourceRef() (string, *string) {
	return "MedicationAdministration", r.ID
}

// This function returns resource reference information
func (r MedicationAdministration) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicationAdministration MedicationAdministration

// MarshalJSON marshals the given MedicationAdministration as JSON into a byte slice
func (r MedicationAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationAdministration
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationAdministration: OtherMedicationAdministration(r),
		ResourceType:                  "MedicationAdministration",
	})
}

// UnmarshalMedicationAdministration unmarshals a MedicationAdministration.
func UnmarshalMedicationAdministration(b []byte) (MedicationAdministration, error) {
	var medicationAdministration MedicationAdministration
	if err := json.Unmarshal(b, &medicationAdministration); err != nil {
		return medicationAdministration, err
	}
	return medicationAdministration, nil
}
