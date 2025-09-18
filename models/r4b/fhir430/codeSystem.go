
package fhir430

import "encoding/json"
// CodeSystem is documented here http://hl7.org/fhir/StructureDefinition/CodeSystem
// The CodeSystem resource is used to declare the existence of and describe a code system or code system supplement and its key properties, and optionally define a part or all of its content.
type CodeSystem struct {
	ID                *string                     `json:"ID,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []json.RawMessage           `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Url               *string                     `json:"url,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Version           *string                     `json:"version,omitempty"`
	Name              *string                     `json:"name,omitempty"`
	Title             *string                     `json:"title,omitempty"`
	Status            PublicationStatus           `json:"status"`
	Experimental      *bool                       `json:"experimental,omitempty"`
	Date              *string                     `json:"date,omitempty"`
	Publisher         *string                     `json:"publisher,omitempty"`
	Contact           []ContactDetail             `json:"contact,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	UseContext        []UsageContext              `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept           `json:"jurisdiction,omitempty"`
	Purpose           *string                     `json:"purpose,omitempty"`
	CaseSensitive     *bool                       `json:"caseSensitive,omitempty"`
	ValueSet          *string                     `json:"valueSet,omitempty"`
	HierarchyMeaning  *CodeSystemHierarchyMeaning `json:"hierarchyMeaning,omitempty"`
	Compositional     *bool                       `json:"compositional,omitempty"`
	VersionNeeded     *bool                       `json:"versionNeeded,omitempty"`
	Content           CodeSystemContentMode       `json:"content"`
	Supplements       *string                     `json:"supplements,omitempty"`
	Count             *int                        `json:"count,omitempty"`
	Filter            []CodeSystemFilter          `json:"filter,omitempty"`
	Property          []CodeSystemProperty        `json:"property,omitempty"`
	Concept           []CodeSystemConcept         `json:"concept,omitempty"`
}

// A filter that can be used in a value set compose statement when selecting concepts using a filter.
// Note that filters defined in code systems usually require custom code on the part of any terminology engine that will make them available for use in value set filters. For this reason, they are generally only seen in high value published terminologies.
type CodeSystemFilter struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              string           `json:"code"`
	Description       *string          `json:"description,omitempty"`
	Operator          []FilterOperator `json:"operator"`
	Value             string           `json:"value"`
}

// A property defines an additional slot through which additional information can be provided about a concept.
type CodeSystemProperty struct {
	ID                *string      `json:"ID,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Code              string       `json:"code"`
	Uri               *string      `json:"uri,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Type              PropertyType `json:"type"`
}

// Concepts that are in the code system. The concept definitions are inherently hierarchical, but the definitions must be consulted to determine what the meanings of the hierarchical relationships are.
// If this is empty, it means that the code system resource does not represent the content of the code system.
type CodeSystemConcept struct {
	ID                *string                        `json:"ID,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              string                         `json:"code"`
	Display           *string                        `json:"display,omitempty"`
	Definition        *string                        `json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `json:"property,omitempty"`
	Concept           []CodeSystemConcept            `json:"concept,omitempty"`
}

// Additional representations for the concept - other languages, aliases, specialized purposes, used for particular purposes, etc.
// Concepts have both a ```display``` and an array of ```designation```. The display is equivalent to a special designation with an implied ```designation.use``` of "primary code" and a language equal to the [Resource Language](resource.html#language).
type CodeSystemConceptDesignation struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Use               *Coding     `json:"use,omitempty"`
	Value             string      `json:"value"`
}

// A property value for this concept.
type CodeSystemConceptProperty struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	ValueCode         string      `json:"valueCode"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueString       string      `json:"valueString"`
	ValueInteger      int         `json:"valueInteger"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueDecimal      json.Number `json:"valueDecimal"`
}

// This function returns resource reference information
func (r CodeSystem) ResourceRef() (string, *string) {
	return "CodeSystem", r.ID
}

// This function returns resource reference information
func (r CodeSystem) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCodeSystem CodeSystem

// MarshalJSON marshals the given CodeSystem as JSON into a byte slice
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}

// UnmarshalCodeSystem unmarshals a CodeSystem.
func UnmarshalCodeSystem(b []byte) (CodeSystem, error) {
	var codeSystem CodeSystem
	if err := json.Unmarshal(b, &codeSystem); err != nil {
		return codeSystem, err
	}
	return codeSystem, nil
}
