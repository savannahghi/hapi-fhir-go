
package fhir430

import "encoding/json"
// NamingSystem is documented here http://hl7.org/fhir/StructureDefinition/NamingSystem
// A curated namespace that issues unique symbols within that namespace for the identification of concepts, people, devices, etc.  Represents a "System" used within the Identifier and Coding data types.
type NamingSystem struct {
	ID                *string                `json:"ID,omitempty"`
	Meta              *Meta                  `json:"meta,omitempty"`
	ImplicitRules     *string                `json:"implicitRules,omitempty"`
	Language          *string                `json:"language,omitempty"`
	Text              *Narrative             `json:"text,omitempty"`
	Contained         []json.RawMessage      `json:"contained,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              string                 `json:"name"`
	Status            PublicationStatus      `json:"status"`
	Kind              NamingSystemType       `json:"kind"`
	Date              string                 `json:"date"`
	Publisher         *string                `json:"publisher,omitempty"`
	Contact           []ContactDetail        `json:"contact,omitempty"`
	Responsible       *string                `json:"responsible,omitempty"`
	Type              *CodeableConcept       `json:"type,omitempty"`
	Description       *string                `json:"description,omitempty"`
	UseContext        []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `json:"jurisdiction,omitempty"`
	Usage             *string                `json:"usage,omitempty"`
	UniqueId          []NamingSystemUniqueId `json:"uniqueId"`
}

// Indicates how the system may be identified when referenced in electronic exchange.
// Multiple identifiers may exist, either due to duplicate registration, regional rules, needs of different communication technologies, etc.
type NamingSystemUniqueId struct {
	ID                *string                    `json:"ID,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Type              NamingSystemIdentifierType `json:"type"`
	Value             string                     `json:"value"`
	Preferred         *bool                      `json:"preferred,omitempty"`
	Comment           *string                    `json:"comment,omitempty"`
	Period            *Period                    `json:"period,omitempty"`
}

// This function returns resource reference information
func (r NamingSystem) ResourceRef() (string, *string) {
	return "NamingSystem", r.ID
}

// This function returns resource reference information
func (r NamingSystem) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherNamingSystem NamingSystem

// MarshalJSON marshals the given NamingSystem as JSON into a byte slice
func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}

// UnmarshalNamingSystem unmarshals a NamingSystem.
func UnmarshalNamingSystem(b []byte) (NamingSystem, error) {
	var namingSystem NamingSystem
	if err := json.Unmarshal(b, &namingSystem); err != nil {
		return namingSystem, err
	}
	return namingSystem, nil
}
