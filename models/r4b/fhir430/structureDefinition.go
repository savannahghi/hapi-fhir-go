package fhir430

import "encoding/json"

// StructureDefinition is documented here http://hl7.org/fhir/StructureDefinition/StructureDefinition
// A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types.
type StructureDefinition struct {
	ID                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Contained         []json.RawMessage                `json:"contained,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Url               string                           `json:"url"`
	Identifier        []Identifier                     `json:"identifier,omitempty"`
	Version           *string                          `json:"version,omitempty"`
	Name              string                           `json:"name"`
	Title             *string                          `json:"title,omitempty"`
	Status            PublicationStatus                `json:"status"`
	Experimental      *bool                            `json:"experimental,omitempty"`
	Date              *string                          `json:"date,omitempty"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose           *string                          `json:"purpose,omitempty"`
	Keyword           []Coding                         `json:"keyword,omitempty"`
	FhirVersion       *FHIRVersion                     `json:"fhirVersion,omitempty"`
	Mapping           []StructureDefinitionMapping     `json:"mapping,omitempty"`
	Kind              StructureDefinitionKind          `json:"kind"`
	Abstract          bool                             `json:"abstract"`
	Context           []StructureDefinitionContext     `json:"context,omitempty"`
	ContextInvariant  []string                         `json:"contextInvariant,omitempty"`
	Type              string                           `json:"type"`
	BaseDefinition    *string                          `json:"baseDefinition,omitempty"`
	Derivation        *TypeDerivationRule              `json:"derivation,omitempty"`
	Snapshot          *StructureDefinitionSnapshot     `json:"snapshot,omitempty"`
	Differential      *StructureDefinitionDifferential `json:"differential,omitempty"`
}

// An external specification that the content is mapped to.
type StructureDefinitionMapping struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identity          string      `json:"identity"`
	Uri               *string     `json:"uri,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
}

// Identifies the types of resource or data type elements to which the extension can be applied.
type StructureDefinitionContext struct {
	ID                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Type              ExtensionContextType `json:"type"`
	Expression        string               `json:"expression"`
}

// A snapshot view is expressed in a standalone form that can be used and interpreted without considering the base StructureDefinition.
type StructureDefinitionSnapshot struct {
	ID                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}

// A differential view is expressed relative to the base StructureDefinition - a statement of differences that it applies.
type StructureDefinitionDifferential struct {
	ID                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}

// This function returns resource reference information
func (r StructureDefinition) ResourceRef() (string, *string) {
	return "StructureDefinition", r.ID
}

// This function returns resource reference information
func (r StructureDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherStructureDefinition StructureDefinition

// MarshalJSON marshals the given StructureDefinition as JSON into a byte slice
func (r StructureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureDefinition: OtherStructureDefinition(r),
		ResourceType:             "StructureDefinition",
	})
}

// UnmarshalStructureDefinition unmarshals a StructureDefinition.
func UnmarshalStructureDefinition(b []byte) (StructureDefinition, error) {
	var structureDefinition StructureDefinition
	if err := json.Unmarshal(b, &structureDefinition); err != nil {
		return structureDefinition, err
	}
	return structureDefinition, nil
}
