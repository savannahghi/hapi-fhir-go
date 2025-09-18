
package fhir430

import "encoding/json"
// VisionPrescription is documented here http://hl7.org/fhir/StructureDefinition/VisionPrescription
// An authorization for the provision of glasses and/or contact lenses to a patient.
type VisionPrescription struct {
	ID                *string                               `json:"ID,omitempty"`
	Meta              *Meta                                 `json:"meta,omitempty"`
	ImplicitRules     *string                               `json:"implicitRules,omitempty"`
	Language          *string                               `json:"language,omitempty"`
	Text              *Narrative                            `json:"text,omitempty"`
	Contained         []json.RawMessage                     `json:"contained,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                          `json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes          `json:"status"`
	Created           string                                `json:"created"`
	Patient           Reference                             `json:"patient"`
	Encounter         *Reference                            `json:"encounter,omitempty"`
	DateWritten       string                                `json:"dateWritten"`
	Prescriber        Reference                             `json:"prescriber"`
	LensSpecification []VisionPrescriptionLensSpecification `json:"lensSpecification"`
}

// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
type VisionPrescriptionLensSpecification struct {
	ID                *string                                    `json:"ID,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Product           CodeableConcept                            `json:"product"`
	Eye               VisionEyes                                 `json:"eye"`
	Sphere            *json.Number                               `json:"sphere,omitempty"`
	Cylinder          *json.Number                               `json:"cylinder,omitempty"`
	Axis              *int                                       `json:"axis,omitempty"`
	Prism             []VisionPrescriptionLensSpecificationPrism `json:"prism,omitempty"`
	Add               *json.Number                               `json:"add,omitempty"`
	Power             *json.Number                               `json:"power,omitempty"`
	BackCurve         *json.Number                               `json:"backCurve,omitempty"`
	Diameter          *json.Number                               `json:"diameter,omitempty"`
	Duration          *Quantity                                  `json:"duration,omitempty"`
	Color             *string                                    `json:"color,omitempty"`
	Brand             *string                                    `json:"brand,omitempty"`
	Note              []Annotation                               `json:"note,omitempty"`
}

// Allows for adjustment on two axis.
type VisionPrescriptionLensSpecificationPrism struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Amount            json.Number `json:"amount"`
	Base              VisionBase  `json:"base"`
}

// This function returns resource reference information
func (r VisionPrescription) ResourceRef() (string, *string) {
	return "VisionPrescription", r.ID
}

// This function returns resource reference information
func (r VisionPrescription) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherVisionPrescription VisionPrescription

// MarshalJSON marshals the given VisionPrescription as JSON into a byte slice
func (r VisionPrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVisionPrescription
		ResourceType string `json:"resourceType"`
	}{
		OtherVisionPrescription: OtherVisionPrescription(r),
		ResourceType:            "VisionPrescription",
	})
}

// UnmarshalVisionPrescription unmarshals a VisionPrescription.
func UnmarshalVisionPrescription(b []byte) (VisionPrescription, error) {
	var visionPrescription VisionPrescription
	if err := json.Unmarshal(b, &visionPrescription); err != nil {
		return visionPrescription, err
	}
	return visionPrescription, nil
}
