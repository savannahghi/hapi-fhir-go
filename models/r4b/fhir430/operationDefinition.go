
package fhir430

import "encoding/json"
// OperationDefinition is documented here http://hl7.org/fhir/StructureDefinition/OperationDefinition
// A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
type OperationDefinition struct {
	ID                *string                        `json:"ID,omitempty"`
	Meta              *Meta                          `json:"meta,omitempty"`
	ImplicitRules     *string                        `json:"implicitRules,omitempty"`
	Language          *string                        `json:"language,omitempty"`
	Text              *Narrative                     `json:"text,omitempty"`
	Contained         []json.RawMessage              `json:"contained,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Url               *string                        `json:"url,omitempty"`
	Version           *string                        `json:"version,omitempty"`
	Name              string                         `json:"name"`
	Title             *string                        `json:"title,omitempty"`
	Status            PublicationStatus              `json:"status"`
	Kind              OperationKind                  `json:"kind"`
	Experimental      *bool                          `json:"experimental,omitempty"`
	Date              *string                        `json:"date,omitempty"`
	Publisher         *string                        `json:"publisher,omitempty"`
	Contact           []ContactDetail                `json:"contact,omitempty"`
	Description       *string                        `json:"description,omitempty"`
	UseContext        []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose           *string                        `json:"purpose,omitempty"`
	AffectsState      *bool                          `json:"affectsState,omitempty"`
	Code              string                         `json:"code"`
	Comment           *string                        `json:"comment,omitempty"`
	Base              *string                        `json:"base,omitempty"`
	Resource          []ResourceType                 `json:"resource,omitempty"`
	System            bool                           `json:"system"`
	Type              bool                           `json:"type"`
	Instance          bool                           `json:"instance"`
	InputProfile      *string                        `json:"inputProfile,omitempty"`
	OutputProfile     *string                        `json:"outputProfile,omitempty"`
	Parameter         []OperationDefinitionParameter `json:"parameter,omitempty"`
	Overload          []OperationDefinitionOverload  `json:"overload,omitempty"`
}

// The parameters for the operation/query.
// Query Definitions only have one output parameter, named "result". This might not be described, but can be to allow a profile to be defined.
type OperationDefinitionParameter struct {
	ID                *string                                      `json:"ID,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Name              string                                       `json:"name"`
	Use               OperationParameterUse                        `json:"use"`
	Min               int                                          `json:"min"`
	Max               string                                       `json:"max"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Type              *string                                      `json:"type,omitempty"`
	TargetProfile     []string                                     `json:"targetProfile,omitempty"`
	SearchType        *SearchParamType                             `json:"searchType,omitempty"`
	Binding           *OperationDefinitionParameterBinding         `json:"binding,omitempty"`
	ReferencedFrom    []OperationDefinitionParameterReferencedFrom `json:"referencedFrom,omitempty"`
	Part              []OperationDefinitionParameter               `json:"part,omitempty"`
}

// Binds to a value set if this parameter is coded (code, Coding, CodeableConcept).
type OperationDefinitionParameterBinding struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Strength          BindingStrength `json:"strength"`
	ValueSet          string          `json:"valueSet"`
}

// Identifies other resource parameters within the operation invocation that are expected to resolve to this resource.
// Resolution applies if the referenced parameter exists.
type OperationDefinitionParameterReferencedFrom struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Source            string      `json:"source"`
	SourceId          *string     `json:"sourceId,omitempty"`
}

// Defines an appropriate combination of parameters to use when invoking this operation, to help code generators when generating overloaded parameter sets for this operation.
// The combinations are suggestions as to which sets of parameters to use together, but the combinations are not intended to be authoritative.
type OperationDefinitionOverload struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ParameterName     []string    `json:"parameterName,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
}

// This function returns resource reference information
func (r OperationDefinition) ResourceRef() (string, *string) {
	return "OperationDefinition", r.ID
}

// This function returns resource reference information
func (r OperationDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherOperationDefinition OperationDefinition

// MarshalJSON marshals the given OperationDefinition as JSON into a byte slice
func (r OperationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationDefinition: OtherOperationDefinition(r),
		ResourceType:             "OperationDefinition",
	})
}

// UnmarshalOperationDefinition unmarshals a OperationDefinition.
func UnmarshalOperationDefinition(b []byte) (OperationDefinition, error) {
	var operationDefinition OperationDefinition
	if err := json.Unmarshal(b, &operationDefinition); err != nil {
		return operationDefinition, err
	}
	return operationDefinition, nil
}
