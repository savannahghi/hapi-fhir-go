
package fhir430

import "encoding/json"
// SubstanceSpecification is documented here http://hl7.org/fhir/StructureDefinition/SubstanceSpecification
// The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceSpecification struct {
	ID                   *string                                                 `json:"ID,omitempty"`
	Meta                 *Meta                                                   `json:"meta,omitempty"`
	ImplicitRules        *string                                                 `json:"implicitRules,omitempty"`
	Language             *string                                                 `json:"language,omitempty"`
	Text                 *Narrative                                              `json:"text,omitempty"`
	Contained            []json.RawMessage                                       `json:"contained,omitempty"`
	Extension            []Extension                                             `json:"extension,omitempty"`
	ModifierExtension    []Extension                                             `json:"modifierExtension,omitempty"`
	Identifier           *Identifier                                             `json:"identifier,omitempty"`
	Type                 *CodeableConcept                                        `json:"type,omitempty"`
	Status               *CodeableConcept                                        `json:"status,omitempty"`
	Domain               *CodeableConcept                                        `json:"domain,omitempty"`
	Description          *string                                                 `json:"description,omitempty"`
	Source               []Reference                                             `json:"source,omitempty"`
	Comment              *string                                                 `json:"comment,omitempty"`
	Moiety               []SubstanceSpecificationMoiety                          `json:"moiety,omitempty"`
	Property             []SubstanceSpecificationProperty                        `json:"property,omitempty"`
	ReferenceInformation *Reference                                              `json:"referenceInformation,omitempty"`
	Structure            *SubstanceSpecificationStructure                        `json:"structure,omitempty"`
	Code                 []SubstanceSpecificationCode                            `json:"code,omitempty"`
	Name                 []SubstanceSpecificationName                            `json:"name,omitempty"`
	MolecularWeight      []SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
	Relationship         []SubstanceSpecificationRelationship                    `json:"relationship,omitempty"`
	NucleicAcid          *Reference                                              `json:"nucleicAcid,omitempty"`
	Polymer              *Reference                                              `json:"polymer,omitempty"`
	Protein              *Reference                                              `json:"protein,omitempty"`
	SourceMaterial       *Reference                                              `json:"sourceMaterial,omitempty"`
}

// Moiety, for structural modifications.
type SubstanceSpecificationMoiety struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Stereochemistry   *CodeableConcept `json:"stereochemistry,omitempty"`
	OpticalActivity   *CodeableConcept `json:"opticalActivity,omitempty"`
	MolecularFormula  *string          `json:"molecularFormula,omitempty"`
	AmountQuantity    *Quantity        `json:"amountQuantity,omitempty"`
	AmountString      *string          `json:"amountString,omitempty"`
}

// General specifications for this substance, including how it is related to other substances.
type SubstanceSpecificationProperty struct {
	ID                               *string          `json:"ID,omitempty"`
	Extension                        []Extension      `json:"extension,omitempty"`
	ModifierExtension                []Extension      `json:"modifierExtension,omitempty"`
	Category                         *CodeableConcept `json:"category,omitempty"`
	Code                             *CodeableConcept `json:"code,omitempty"`
	Parameters                       *string          `json:"parameters,omitempty"`
	DefiningSubstanceReference       *Reference       `json:"definingSubstanceReference,omitempty"`
	DefiningSubstanceCodeableConcept *CodeableConcept `json:"definingSubstanceCodeableConcept,omitempty"`
	AmountQuantity                   *Quantity        `json:"amountQuantity,omitempty"`
	AmountString                     *string          `json:"amountString,omitempty"`
}

// Structural information.
type SubstanceSpecificationStructure struct {
	ID                       *string                                                `json:"ID,omitempty"`
	Extension                []Extension                                            `json:"extension,omitempty"`
	ModifierExtension        []Extension                                            `json:"modifierExtension,omitempty"`
	Stereochemistry          *CodeableConcept                                       `json:"stereochemistry,omitempty"`
	OpticalActivity          *CodeableConcept                                       `json:"opticalActivity,omitempty"`
	MolecularFormula         *string                                                `json:"molecularFormula,omitempty"`
	MolecularFormulaByMoiety *string                                                `json:"molecularFormulaByMoiety,omitempty"`
	Isotope                  []SubstanceSpecificationStructureIsotope               `json:"isotope,omitempty"`
	MolecularWeight          *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
	Source                   []Reference                                            `json:"source,omitempty"`
	Representation           []SubstanceSpecificationStructureRepresentation        `json:"representation,omitempty"`
}

// Applicable for single substances that contain a radionuclide or a non-natural isotopic ratio.
type SubstanceSpecificationStructureIsotope struct {
	ID                *string                                                `json:"ID,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                                            `json:"identifier,omitempty"`
	Name              *CodeableConcept                                       `json:"name,omitempty"`
	Substitution      *CodeableConcept                                       `json:"substitution,omitempty"`
	HalfLife          *Quantity                                              `json:"halfLife,omitempty"`
	MolecularWeight   *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
}

// The molecular weight or weight range (for proteins, polymers or nucleic acids).
type SubstanceSpecificationStructureIsotopeMolecularWeight struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

// Molecular structural representation.
type SubstanceSpecificationStructureRepresentation struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Attachment        *Attachment      `json:"attachment,omitempty"`
}

// Codes associated with the substance.
type SubstanceSpecificationCode struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	StatusDate        *string          `json:"statusDate,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

// Names applicable to this substance.
type SubstanceSpecificationName struct {
	ID                *string                              `json:"ID,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Name              string                               `json:"name"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	Status            *CodeableConcept                     `json:"status,omitempty"`
	Preferred         *bool                                `json:"preferred,omitempty"`
	Language          []CodeableConcept                    `json:"language,omitempty"`
	Domain            []CodeableConcept                    `json:"domain,omitempty"`
	Jurisdiction      []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Synonym           []SubstanceSpecificationName         `json:"synonym,omitempty"`
	Translation       []SubstanceSpecificationName         `json:"translation,omitempty"`
	Official          []SubstanceSpecificationNameOfficial `json:"official,omitempty"`
	Source            []Reference                          `json:"source,omitempty"`
}

// Details of the official nature of this name.
type SubstanceSpecificationNameOfficial struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `json:"authority,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	Date              *string          `json:"date,omitempty"`
}

// A link between this substance and another, with details of the relationship.
type SubstanceSpecificationRelationship struct {
	ID                       *string          `json:"ID,omitempty"`
	Extension                []Extension      `json:"extension,omitempty"`
	ModifierExtension        []Extension      `json:"modifierExtension,omitempty"`
	SubstanceReference       *Reference       `json:"substanceReference,omitempty"`
	SubstanceCodeableConcept *CodeableConcept `json:"substanceCodeableConcept,omitempty"`
	Relationship             *CodeableConcept `json:"relationship,omitempty"`
	IsDefining               *bool            `json:"isDefining,omitempty"`
	AmountQuantity           *Quantity        `json:"amountQuantity,omitempty"`
	AmountRange              *Range           `json:"amountRange,omitempty"`
	AmountRatio              *Ratio           `json:"amountRatio,omitempty"`
	AmountString             *string          `json:"amountString,omitempty"`
	AmountRatioLowLimit      *Ratio           `json:"amountRatioLowLimit,omitempty"`
	AmountType               *CodeableConcept `json:"amountType,omitempty"`
	Source                   []Reference      `json:"source,omitempty"`
}

// This function returns resource reference information
func (r SubstanceSpecification) ResourceRef() (string, *string) {
	return "SubstanceSpecification", r.ID
}

// This function returns resource reference information
func (r SubstanceSpecification) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSubstanceSpecification SubstanceSpecification

// MarshalJSON marshals the given SubstanceSpecification as JSON into a byte slice
func (r SubstanceSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceSpecification
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceSpecification: OtherSubstanceSpecification(r),
		ResourceType:                "SubstanceSpecification",
	})
}

// UnmarshalSubstanceSpecification unmarshals a SubstanceSpecification.
func UnmarshalSubstanceSpecification(b []byte) (SubstanceSpecification, error) {
	var substanceSpecification SubstanceSpecification
	if err := json.Unmarshal(b, &substanceSpecification); err != nil {
		return substanceSpecification, err
	}
	return substanceSpecification, nil
}
