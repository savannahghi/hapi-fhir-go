
package fhir430

import "encoding/json"
// CompartmentDefinition is documented here http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
// A compartment definition that defines how resources are accessed on a server.
type CompartmentDefinition struct {
	ID                *string                         `json:"ID,omitempty"`
	Meta              *Meta                           `json:"meta,omitempty"`
	ImplicitRules     *string                         `json:"implicitRules,omitempty"`
	Language          *string                         `json:"language,omitempty"`
	Text              *Narrative                      `json:"text,omitempty"`
	Contained         []json.RawMessage               `json:"contained,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Url               string                          `json:"url"`
	Version           *string                         `json:"version,omitempty"`
	Name              string                          `json:"name"`
	Status            PublicationStatus               `json:"status"`
	Experimental      *bool                           `json:"experimental,omitempty"`
	Date              *string                         `json:"date,omitempty"`
	Publisher         *string                         `json:"publisher,omitempty"`
	Contact           []ContactDetail                 `json:"contact,omitempty"`
	Description       *string                         `json:"description,omitempty"`
	UseContext        []UsageContext                  `json:"useContext,omitempty"`
	Purpose           *string                         `json:"purpose,omitempty"`
	Code              CompartmentType                 `json:"code"`
	Search            bool                            `json:"search"`
	Resource          []CompartmentDefinitionResource `json:"resource,omitempty"`
}

// Information about how a resource is related to the compartment.
type CompartmentDefinitionResource struct {
	ID                *string      `json:"ID,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Code              ResourceType `json:"code"`
	Param             []string     `json:"param,omitempty"`
	Documentation     *string      `json:"documentation,omitempty"`
}

// This function returns resource reference information
func (r CompartmentDefinition) ResourceRef() (string, *string) {
	return "CompartmentDefinition", r.ID
}

// This function returns resource reference information
func (r CompartmentDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCompartmentDefinition CompartmentDefinition

// MarshalJSON marshals the given CompartmentDefinition as JSON into a byte slice
func (r CompartmentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCompartmentDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherCompartmentDefinition: OtherCompartmentDefinition(r),
		ResourceType:               "CompartmentDefinition",
	})
}

// UnmarshalCompartmentDefinition unmarshals a CompartmentDefinition.
func UnmarshalCompartmentDefinition(b []byte) (CompartmentDefinition, error) {
	var compartmentDefinition CompartmentDefinition
	if err := json.Unmarshal(b, &compartmentDefinition); err != nil {
		return compartmentDefinition, err
	}
	return compartmentDefinition, nil
}
