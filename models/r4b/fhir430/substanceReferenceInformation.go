package fhir430

import "encoding/json"

// SubstanceReferenceInformation is documented here http://hl7.org/fhir/StructureDefinition/SubstanceReferenceInformation
// Todo.
type SubstanceReferenceInformation struct {
	ID                *string                                       `json:"id,omitempty"`
	Meta              *Meta                                         `json:"meta,omitempty"`
	ImplicitRules     *string                                       `json:"implicitRules,omitempty"`
	Language          *string                                       `json:"language,omitempty"`
	Text              *Narrative                                    `json:"text,omitempty"`
	Contained         []json.RawMessage                             `json:"contained,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Comment           *string                                       `json:"comment,omitempty"`
	Gene              []SubstanceReferenceInformationGene           `json:"gene,omitempty"`
	GeneElement       []SubstanceReferenceInformationGeneElement    `json:"geneElement,omitempty"`
	Classification    []SubstanceReferenceInformationClassification `json:"classification,omitempty"`
	Target            []SubstanceReferenceInformationTarget         `json:"target,omitempty"`
}

// Todo.
type SubstanceReferenceInformationGene struct {
	ID                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	GeneSequenceOrigin *CodeableConcept `json:"geneSequenceOrigin,omitempty"`
	Gene               *CodeableConcept `json:"gene,omitempty"`
	Source             []Reference      `json:"source,omitempty"`
}

// Todo.
type SubstanceReferenceInformationGeneElement struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Element           *Identifier      `json:"element,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

// Todo.
type SubstanceReferenceInformationClassification struct {
	ID                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Domain            *CodeableConcept  `json:"domain,omitempty"`
	Classification    *CodeableConcept  `json:"classification,omitempty"`
	Subtype           []CodeableConcept `json:"subtype,omitempty"`
	Source            []Reference       `json:"source,omitempty"`
}

// Todo.
type SubstanceReferenceInformationTarget struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Target            *Identifier      `json:"target,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Interaction       *CodeableConcept `json:"interaction,omitempty"`
	Organism          *CodeableConcept `json:"organism,omitempty"`
	OrganismType      *CodeableConcept `json:"organismType,omitempty"`
	AmountQuantity    *Quantity        `json:"amountQuantity,omitempty"`
	AmountRange       *Range           `json:"amountRange,omitempty"`
	AmountString      *string          `json:"amountString,omitempty"`
	AmountType        *CodeableConcept `json:"amountType,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

// This function returns resource reference information
func (r SubstanceReferenceInformation) ResourceRef() (string, *string) {
	return "SubstanceReferenceInformation", r.ID
}

// This function returns resource reference information
func (r SubstanceReferenceInformation) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSubstanceReferenceInformation SubstanceReferenceInformation

// MarshalJSON marshals the given SubstanceReferenceInformation as JSON into a byte slice
func (r SubstanceReferenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceReferenceInformation
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceReferenceInformation: OtherSubstanceReferenceInformation(r),
		ResourceType:                       "SubstanceReferenceInformation",
	})
}

// UnmarshalSubstanceReferenceInformation unmarshals a SubstanceReferenceInformation.
func UnmarshalSubstanceReferenceInformation(b []byte) (SubstanceReferenceInformation, error) {
	var substanceReferenceInformation SubstanceReferenceInformation
	if err := json.Unmarshal(b, &substanceReferenceInformation); err != nil {
		return substanceReferenceInformation, err
	}
	return substanceReferenceInformation, nil
}
