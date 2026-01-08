package fhir430

import "encoding/json"

// GraphDefinition is documented here http://hl7.org/fhir/StructureDefinition/GraphDefinition
// A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinition struct {
	ID                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []json.RawMessage     `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Url               *string               `json:"url,omitempty"`
	Version           *string               `json:"version,omitempty"`
	Name              string                `json:"name"`
	Status            PublicationStatus     `json:"status"`
	Experimental      *bool                 `json:"experimental,omitempty"`
	Date              *string               `json:"date,omitempty"`
	Publisher         *string               `json:"publisher,omitempty"`
	Contact           []ContactDetail       `json:"contact,omitempty"`
	Description       *string               `json:"description,omitempty"`
	UseContext        []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose           *string               `json:"purpose,omitempty"`
	Start             ResourceType          `json:"start"`
	Profile           *string               `json:"profile,omitempty"`
	Link              []GraphDefinitionLink `json:"link,omitempty"`
}

// Links this graph makes rules about.
type GraphDefinitionLink struct {
	ID                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Path              *string                     `json:"path,omitempty"`
	SliceName         *string                     `json:"sliceName,omitempty"`
	Min               *int                        `json:"min,omitempty"`
	Max               *string                     `json:"max,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Target            []GraphDefinitionLinkTarget `json:"target,omitempty"`
}

// Potential target for the link.
type GraphDefinitionLinkTarget struct {
	ID                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Type              ResourceType                           `json:"type"`
	Params            *string                                `json:"params,omitempty"`
	Profile           *string                                `json:"profile,omitempty"`
	Compartment       []GraphDefinitionLinkTargetCompartment `json:"compartment,omitempty"`
	Link              []GraphDefinitionLink                  `json:"link,omitempty"`
}

// Compartment Consistency Rules.
type GraphDefinitionLinkTargetCompartment struct {
	ID                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Use               GraphCompartmentUse  `json:"use"`
	Code              CompartmentType      `json:"code"`
	Rule              GraphCompartmentRule `json:"rule"`
	Expression        *string              `json:"expression,omitempty"`
	Description       *string              `json:"description,omitempty"`
}

// This function returns resource reference information
func (r GraphDefinition) ResourceRef() (string, *string) {
	return "GraphDefinition", r.ID
}

// This function returns resource reference information
func (r GraphDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherGraphDefinition GraphDefinition

// MarshalJSON marshals the given GraphDefinition as JSON into a byte slice
func (r GraphDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
}

// UnmarshalGraphDefinition unmarshals a GraphDefinition.
func UnmarshalGraphDefinition(b []byte) (GraphDefinition, error) {
	var graphDefinition GraphDefinition
	if err := json.Unmarshal(b, &graphDefinition); err != nil {
		return graphDefinition, err
	}
	return graphDefinition, nil
}
