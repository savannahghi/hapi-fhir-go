
package fhir430

import "encoding/json"
// MedicinalProductIndication is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductIndication
// Indication for the Medicinal Product.
type MedicinalProductIndication struct {
	ID                      *string                                  `json:"ID,omitempty"`
	Meta                    *Meta                                    `json:"meta,omitempty"`
	ImplicitRules           *string                                  `json:"implicitRules,omitempty"`
	Language                *string                                  `json:"language,omitempty"`
	Text                    *Narrative                               `json:"text,omitempty"`
	Contained               []json.RawMessage                        `json:"contained,omitempty"`
	Extension               []Extension                              `json:"extension,omitempty"`
	ModifierExtension       []Extension                              `json:"modifierExtension,omitempty"`
	Subject                 []Reference                              `json:"subject,omitempty"`
	DiseaseSymptomProcedure *CodeableConcept                         `json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus           *CodeableConcept                         `json:"diseaseStatus,omitempty"`
	Comorbidity             []CodeableConcept                        `json:"comorbidity,omitempty"`
	IntendedEffect          *CodeableConcept                         `json:"intendedEffect,omitempty"`
	Duration                *Quantity                                `json:"duration,omitempty"`
	OtherTherapy            []MedicinalProductIndicationOtherTherapy `json:"otherTherapy,omitempty"`
	UndesirableEffect       []Reference                              `json:"undesirableEffect,omitempty"`
	Population              []Population                             `json:"population,omitempty"`
}

// Information about the use of the medicinal product in relation to other therapies described as part of the indication.
type MedicinalProductIndicationOtherTherapy struct {
	ID                        *string         `json:"ID,omitempty"`
	Extension                 []Extension     `json:"extension,omitempty"`
	ModifierExtension         []Extension     `json:"modifierExtension,omitempty"`
	TherapyRelationshipType   CodeableConcept `json:"therapyRelationshipType"`
	MedicationCodeableConcept CodeableConcept `json:"medicationCodeableConcept"`
	MedicationReference       Reference       `json:"medicationReference"`
}

// This function returns resource reference information
func (r MedicinalProductIndication) ResourceRef() (string, *string) {
	return "MedicinalProductIndication", r.ID
}

// This function returns resource reference information
func (r MedicinalProductIndication) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductIndication MedicinalProductIndication

// MarshalJSON marshals the given MedicinalProductIndication as JSON into a byte slice
func (r MedicinalProductIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductIndication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductIndication: OtherMedicinalProductIndication(r),
		ResourceType:                    "MedicinalProductIndication",
	})
}

// UnmarshalMedicinalProductIndication unmarshals a MedicinalProductIndication.
func UnmarshalMedicinalProductIndication(b []byte) (MedicinalProductIndication, error) {
	var medicinalProductIndication MedicinalProductIndication
	if err := json.Unmarshal(b, &medicinalProductIndication); err != nil {
		return medicinalProductIndication, err
	}
	return medicinalProductIndication, nil
}
