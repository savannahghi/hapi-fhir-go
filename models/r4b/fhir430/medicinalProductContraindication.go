package fhir430

import "encoding/json"

// MedicinalProductContraindication is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductContraindication
// The clinical particulars - indications, contraindications etc. of a medicinal product, including for regulatory purposes.
type MedicinalProductContraindication struct {
	ID                    *string                                        `json:"id,omitempty"`
	Meta                  *Meta                                          `json:"meta,omitempty"`
	ImplicitRules         *string                                        `json:"implicitRules,omitempty"`
	Language              *string                                        `json:"language,omitempty"`
	Text                  *Narrative                                     `json:"text,omitempty"`
	Contained             []json.RawMessage                              `json:"contained,omitempty"`
	Extension             []Extension                                    `json:"extension,omitempty"`
	ModifierExtension     []Extension                                    `json:"modifierExtension,omitempty"`
	Subject               []Reference                                    `json:"subject,omitempty"`
	Disease               *CodeableConcept                               `json:"disease,omitempty"`
	DiseaseStatus         *CodeableConcept                               `json:"diseaseStatus,omitempty"`
	Comorbidity           []CodeableConcept                              `json:"comorbidity,omitempty"`
	TherapeuticIndication []Reference                                    `json:"therapeuticIndication,omitempty"`
	OtherTherapy          []MedicinalProductContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
	Population            []Population                                   `json:"population,omitempty"`
}

// Information about the use of the medicinal product in relation to other therapies described as part of the indication.
type MedicinalProductContraindicationOtherTherapy struct {
	ID                        *string         `json:"id,omitempty"`
	Extension                 []Extension     `json:"extension,omitempty"`
	ModifierExtension         []Extension     `json:"modifierExtension,omitempty"`
	TherapyRelationshipType   CodeableConcept `json:"therapyRelationshipType"`
	MedicationCodeableConcept CodeableConcept `json:"medicationCodeableConcept"`
	MedicationReference       Reference       `json:"medicationReference"`
}

// This function returns resource reference information
func (r MedicinalProductContraindication) ResourceRef() (string, *string) {
	return "MedicinalProductContraindication", r.ID
}

// This function returns resource reference information
func (r MedicinalProductContraindication) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductContraindication MedicinalProductContraindication

// MarshalJSON marshals the given MedicinalProductContraindication as JSON into a byte slice
func (r MedicinalProductContraindication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductContraindication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductContraindication: OtherMedicinalProductContraindication(r),
		ResourceType:                          "MedicinalProductContraindication",
	})
}

// UnmarshalMedicinalProductContraindication unmarshals a MedicinalProductContraindication.
func UnmarshalMedicinalProductContraindication(b []byte) (MedicinalProductContraindication, error) {
	var medicinalProductContraindication MedicinalProductContraindication
	if err := json.Unmarshal(b, &medicinalProductContraindication); err != nil {
		return medicinalProductContraindication, err
	}
	return medicinalProductContraindication, nil
}
