package fhir430

import "encoding/json"

// ConceptMap is documented here http://hl7.org/fhir/StructureDefinition/ConceptMap
// A statement of relationships from one set of concepts to one or more other concepts - either concepts in code systems, or data element/data element concepts, or classes in class models.
type ConceptMap struct {
	ID                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Url               *string           `json:"url,omitempty"`
	Identifier        *Identifier       `json:"identifier,omitempty"`
	Version           *string           `json:"version,omitempty"`
	Name              *string           `json:"name,omitempty"`
	Title             *string           `json:"title,omitempty"`
	Status            PublicationStatus `json:"status"`
	Experimental      *bool             `json:"experimental,omitempty"`
	Date              *string           `json:"date,omitempty"`
	Publisher         *string           `json:"publisher,omitempty"`
	Contact           []ContactDetail   `json:"contact,omitempty"`
	Description       *string           `json:"description,omitempty"`
	UseContext        []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose           *string           `json:"purpose,omitempty"`
	SourceUri         *string           `json:"sourceUri,omitempty"`
	SourceCanonical   *string           `json:"sourceCanonical,omitempty"`
	TargetUri         *string           `json:"targetUri,omitempty"`
	TargetCanonical   *string           `json:"targetCanonical,omitempty"`
	Group             []ConceptMapGroup `json:"group,omitempty"`
}

// A group of mappings that all have the same source and target system.
type ConceptMapGroup struct {
	ID                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Source            *string                  `json:"source,omitempty"`
	SourceVersion     *string                  `json:"sourceVersion,omitempty"`
	Target            *string                  `json:"target,omitempty"`
	TargetVersion     *string                  `json:"targetVersion,omitempty"`
	Element           []ConceptMapGroupElement `json:"element"`
	Unmapped          *ConceptMapGroupUnmapped `json:"unmapped,omitempty"`
}

// Mappings for an individual concept in the source to one or more concepts in the target.
// Generally, the ideal is that there would only be one mapping for each concept in the source value set, but a given concept may be mapped multiple times with different comments or dependencies.
type ConceptMapGroupElement struct {
	ID                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *string                        `json:"code,omitempty"`
	Display           *string                        `json:"display,omitempty"`
	Target            []ConceptMapGroupElementTarget `json:"target,omitempty"`
}

// A concept from the target value set that this concept maps to.
// Ideally there would only be one map, with equal or equivalent mapping. But multiple maps are allowed for several narrower options, or to assert that other concepts are unmatched.
type ConceptMapGroupElementTarget struct {
	ID                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Code              *string                                 `json:"code,omitempty"`
	Display           *string                                 `json:"display,omitempty"`
	Equivalence       ConceptMapEquivalence                   `json:"equivalence"`
	Comment           *string                                 `json:"comment,omitempty"`
	DependsOn         []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty"`
	Product           []ConceptMapGroupElementTargetDependsOn `json:"product,omitempty"`
}

// A set of additional dependencies for this mapping to hold. This mapping is only applicable if the specified element can be resolved, and it has the specified value.
type ConceptMapGroupElementTargetDependsOn struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Property          string      `json:"property"`
	System            *string     `json:"system,omitempty"`
	Value             string      `json:"value"`
	Display           *string     `json:"display,omitempty"`
}

// What to do when there is no mapping for the source concept. "Unmapped" does not include codes that are unmatched, and the unmapped element is ignored in a code is specified to have equivalence = unmatched.
// This only applies if the source code has a system value that matches the system defined for the group.
type ConceptMapGroupUnmapped struct {
	ID                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Mode              ConceptMapGroupUnmappedMode `json:"mode"`
	Code              *string                     `json:"code,omitempty"`
	Display           *string                     `json:"display,omitempty"`
	Url               *string                     `json:"url,omitempty"`
}

// This function returns resource reference information
func (r ConceptMap) ResourceRef() (string, *string) {
	return "ConceptMap", r.ID
}

// This function returns resource reference information
func (r ConceptMap) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherConceptMap ConceptMap

// MarshalJSON marshals the given ConceptMap as JSON into a byte slice
func (r ConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConceptMap
		ResourceType string `json:"resourceType"`
	}{
		OtherConceptMap: OtherConceptMap(r),
		ResourceType:    "ConceptMap",
	})
}

// UnmarshalConceptMap unmarshals a ConceptMap.
func UnmarshalConceptMap(b []byte) (ConceptMap, error) {
	var conceptMap ConceptMap
	if err := json.Unmarshal(b, &conceptMap); err != nil {
		return conceptMap, err
	}
	return conceptMap, nil
}
