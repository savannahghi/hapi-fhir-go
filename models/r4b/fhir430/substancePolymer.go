package fhir430

import "encoding/json"

// SubstancePolymer is documented here http://hl7.org/fhir/StructureDefinition/SubstancePolymer
// Todo.
type SubstancePolymer struct {
	ID                    *string                      `json:"id,omitempty"`
	Meta                  *Meta                        `json:"meta,omitempty"`
	ImplicitRules         *string                      `json:"implicitRules,omitempty"`
	Language              *string                      `json:"language,omitempty"`
	Text                  *Narrative                   `json:"text,omitempty"`
	Contained             []json.RawMessage            `json:"contained,omitempty"`
	Extension             []Extension                  `json:"extension,omitempty"`
	ModifierExtension     []Extension                  `json:"modifierExtension,omitempty"`
	Class                 *CodeableConcept             `json:"class,omitempty"`
	Geometry              *CodeableConcept             `json:"geometry,omitempty"`
	CopolymerConnectivity []CodeableConcept            `json:"copolymerConnectivity,omitempty"`
	Modification          []string                     `json:"modification,omitempty"`
	MonomerSet            []SubstancePolymerMonomerSet `json:"monomerSet,omitempty"`
	Repeat                []SubstancePolymerRepeat     `json:"repeat,omitempty"`
}

// Todo.
type SubstancePolymerMonomerSet struct {
	ID                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	RatioType         *CodeableConcept                             `json:"ratioType,omitempty"`
	StartingMaterial  []SubstancePolymerMonomerSetStartingMaterial `json:"startingMaterial,omitempty"`
}

// Todo.
type SubstancePolymerMonomerSetStartingMaterial struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Material          *CodeableConcept `json:"material,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	IsDefining        *bool            `json:"isDefining,omitempty"`
	Amount            *SubstanceAmount `json:"amount,omitempty"`
}

// Todo.
type SubstancePolymerRepeat struct {
	ID                      *string                            `json:"id,omitempty"`
	Extension               []Extension                        `json:"extension,omitempty"`
	ModifierExtension       []Extension                        `json:"modifierExtension,omitempty"`
	NumberOfUnits           *int                               `json:"numberOfUnits,omitempty"`
	AverageMolecularFormula *string                            `json:"averageMolecularFormula,omitempty"`
	RepeatUnitAmountType    *CodeableConcept                   `json:"repeatUnitAmountType,omitempty"`
	RepeatUnit              []SubstancePolymerRepeatRepeatUnit `json:"repeatUnit,omitempty"`
}

// Todo.
type SubstancePolymerRepeatRepeatUnit struct {
	ID                          *string                                                    `json:"id,omitempty"`
	Extension                   []Extension                                                `json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `json:"modifierExtension,omitempty"`
	OrientationOfPolymerisation *CodeableConcept                                           `json:"orientationOfPolymerisation,omitempty"`
	RepeatUnit                  *string                                                    `json:"repeatUnit,omitempty"`
	Amount                      *SubstanceAmount                                           `json:"amount,omitempty"`
	DegreeOfPolymerisation      []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation   `json:"degreeOfPolymerisation,omitempty"`
	StructuralRepresentation    []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `json:"structuralRepresentation,omitempty"`
}

// Todo.
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Degree            *CodeableConcept `json:"degree,omitempty"`
	Amount            *SubstanceAmount `json:"amount,omitempty"`
}

// Todo.
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Attachment        *Attachment      `json:"attachment,omitempty"`
}

// This function returns resource reference information
func (r SubstancePolymer) ResourceRef() (string, *string) {
	return "SubstancePolymer", r.ID
}

// This function returns resource reference information
func (r SubstancePolymer) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSubstancePolymer SubstancePolymer

// MarshalJSON marshals the given SubstancePolymer as JSON into a byte slice
func (r SubstancePolymer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstancePolymer
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstancePolymer: OtherSubstancePolymer(r),
		ResourceType:          "SubstancePolymer",
	})
}

// UnmarshalSubstancePolymer unmarshals a SubstancePolymer.
func UnmarshalSubstancePolymer(b []byte) (SubstancePolymer, error) {
	var substancePolymer SubstancePolymer
	if err := json.Unmarshal(b, &substancePolymer); err != nil {
		return substancePolymer, err
	}
	return substancePolymer, nil
}
