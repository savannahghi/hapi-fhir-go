package models

import "encoding/json"

// SubstanceDefinition is documented here http://hl7.org/fhir/StructureDefinition/SubstanceDefinition
type SubstanceDefinition struct {
	ID                   *string                               `json:"id,omitempty"`
	Meta                 *Meta                                 `json:"meta,omitempty"`
	ImplicitRules        *string                               `json:"implicitRules,omitempty"`
	Language             *string                               `json:"language,omitempty"`
	Text                 *Narrative                            `json:"text,omitempty"`
	Extension            []Extension                           `json:"extension,omitempty"`
	ModifierExtension    []Extension                           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                          `json:"identifier,omitempty"`
	Version              *string                               `json:"version,omitempty"`
	Status               *CodeableConcept                      `json:"status,omitempty"`
	Classification       []CodeableConcept                     `json:"classification,omitempty"`
	Domain               *CodeableConcept                      `json:"domain,omitempty"`
	Grade                []CodeableConcept                     `json:"grade,omitempty"`
	Description          *string                               `json:"description,omitempty"`
	InformationSource    []Reference                           `json:"informationSource,omitempty"`
	Note                 []Annotation                          `json:"note,omitempty"`
	Manufacturer         []Reference                           `json:"manufacturer,omitempty"`
	Supplier             []Reference                           `json:"supplier,omitempty"`
	Moiety               []SubstanceDefinitionMoiety           `json:"moiety,omitempty"`
	Characterization     []SubstanceDefinitionCharacterization `json:"characterization,omitempty"`
	Property             []SubstanceDefinitionProperty         `json:"property,omitempty"`
	ReferenceInformation *Reference                            `json:"referenceInformation,omitempty"`
	MolecularWeight      []SubstanceDefinitionMolecularWeight  `json:"molecularWeight,omitempty"`
	Structure            *SubstanceDefinitionStructure         `json:"structure,omitempty"`
	Code                 []SubstanceDefinitionCode             `json:"code,omitempty"`
	Name                 []SubstanceDefinitionName             `json:"name,omitempty"`
	Relationship         []SubstanceDefinitionRelationship     `json:"relationship,omitempty"`
	NucleicAcid          *Reference                            `json:"nucleicAcid,omitempty"`
	Polymer              *Reference                            `json:"polymer,omitempty"`
	Protein              *Reference                            `json:"protein,omitempty"`
	SourceMaterial       *SubstanceDefinitionSourceMaterial    `json:"sourceMaterial,omitempty"`
}

type SubstanceDefinitionMoiety struct {
	ID                *string          `json:"id,omitempty"`
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
	MeasurementType   *CodeableConcept `json:"measurementType,omitempty"`
}

type SubstanceDefinitionCharacterization struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Technique         *CodeableConcept `json:"technique,omitempty"`
	Form              *CodeableConcept `json:"form,omitempty"`
	Description       *string          `json:"description,omitempty"`
	File              []Attachment     `json:"file,omitempty"`
}

type SubstanceDefinitionProperty struct {
	ID                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string          `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
}

type SubstanceDefinitionMolecularWeight struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Amount            Quantity         `json:"amount"`
}

type SubstanceDefinitionStructure struct {
	ID                       *string                                      `json:"id,omitempty"`
	Extension                []Extension                                  `json:"extension,omitempty"`
	ModifierExtension        []Extension                                  `json:"modifierExtension,omitempty"`
	Stereochemistry          *CodeableConcept                             `json:"stereochemistry,omitempty"`
	OpticalActivity          *CodeableConcept                             `json:"opticalActivity,omitempty"`
	MolecularFormula         *string                                      `json:"molecularFormula,omitempty"`
	MolecularFormulaByMoiety *string                                      `json:"molecularFormulaByMoiety,omitempty"`
	MolecularWeight          *SubstanceDefinitionMolecularWeight          `json:"molecularWeight,omitempty"`
	Technique                []CodeableConcept                            `json:"technique,omitempty"`
	SourceDocument           []Reference                                  `json:"sourceDocument,omitempty"`
	Representation           []SubstanceDefinitionStructureRepresentation `json:"representation,omitempty"`
}

type SubstanceDefinitionStructureRepresentation struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Format            *CodeableConcept `json:"format,omitempty"`
	Document          *Reference       `json:"document,omitempty"`
}

type SubstanceDefinitionCode struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	StatusDate        *string          `json:"statusDate,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

type SubstanceDefinitionName struct {
	ID                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Name              string                            `json:"name"`
	Type              *CodeableConcept                  `json:"type,omitempty"`
	Status            *CodeableConcept                  `json:"status,omitempty"`
	Preferred         *bool                             `json:"preferred,omitempty"`
	Language          []CodeableConcept                 `json:"language,omitempty"`
	Domain            []CodeableConcept                 `json:"domain,omitempty"`
	Jurisdiction      []CodeableConcept                 `json:"jurisdiction,omitempty"`
	Synonym           []SubstanceDefinitionName         `json:"synonym,omitempty"`
	Translation       []SubstanceDefinitionName         `json:"translation,omitempty"`
	Official          []SubstanceDefinitionNameOfficial `json:"official,omitempty"`
	Source            []Reference                       `json:"source,omitempty"`
}

type SubstanceDefinitionNameOfficial struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `json:"authority,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	Date              *string          `json:"date,omitempty"`
}

type SubstanceDefinitionRelationship struct {
	ID                                 *string          `json:"id,omitempty"`
	Extension                          []Extension      `json:"extension,omitempty"`
	ModifierExtension                  []Extension      `json:"modifierExtension,omitempty"`
	SubstanceDefinitionReference       *Reference       `json:"substanceDefinitionReference,omitempty"`
	SubstanceDefinitionCodeableConcept *CodeableConcept `json:"substanceDefinitionCodeableConcept,omitempty"`
	Type                               CodeableConcept  `json:"type"`
	IsDefining                         *bool            `json:"isDefining,omitempty"`
	AmountQuantity                     *Quantity        `json:"amountQuantity,omitempty"`
	AmountRatio                        *Ratio           `json:"amountRatio,omitempty"`
	AmountString                       *string          `json:"amountString,omitempty"`
	RatioHighLimitAmount               *Ratio           `json:"ratioHighLimitAmount,omitempty"`
	Comparator                         *CodeableConcept `json:"comparator,omitempty"`
	Source                             []Reference      `json:"source,omitempty"`
}

type SubstanceDefinitionSourceMaterial struct {
	ID                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Genus             *CodeableConcept  `json:"genus,omitempty"`
	Species           *CodeableConcept  `json:"species,omitempty"`
	Part              *CodeableConcept  `json:"part,omitempty"`
	CountryOfOrigin   []CodeableConcept `json:"countryOfOrigin,omitempty"`
}

type OtherSubstanceDefinition SubstanceDefinition

// MarshalJSON marshals the given SubstanceDefinition as JSON into a byte slice.
func (r SubstanceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceDefinition: OtherSubstanceDefinition(r),
		ResourceType:             "SubstanceDefinition",
	})
}

// UnmarshalSubstanceDefinition unmarshals a SubstanceDefinition.
func UnmarshalSubstanceDefinition(b []byte) (SubstanceDefinition, error) {
	var substanceDefinition SubstanceDefinition
	if err := json.Unmarshal(b, &substanceDefinition); err != nil {
		return substanceDefinition, err
	}

	return substanceDefinition, nil
}
