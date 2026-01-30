package fhir430

import "encoding/json"

// Group is documented here http://hl7.org/fhir/StructureDefinition/Group
// Represents a defined collection of entities that may be discussed or acted upon collectively but which are not expected to act collectively, and are not formally or legally recognized; i.e. a collection of entities that isn't an Organization.
type Group struct {
	ID                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []json.RawMessage     `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Active            *bool                 `json:"active,omitempty"`
	Type              GroupType             `json:"type"`
	Actual            bool                  `json:"actual"`
	Code              *CodeableConcept      `json:"code,omitempty"`
	Name              *string               `json:"name,omitempty"`
	Quantity          *int                  `json:"quantity,omitempty"`
	ManagingEntity    *Reference            `json:"managingEntity,omitempty"`
	Characteristic    []GroupCharacteristic `json:"characteristic,omitempty"`
	Member            []GroupMember         `json:"member,omitempty"`
}

// Identifies traits whose presence r absence is shared by members of the group.
// All the identified characteristics must be true for an entity to a member of the group.
type GroupCharacteristic struct {
	ID                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Code                 CodeableConcept `json:"code"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueReference       Reference       `json:"valueReference"`
	Exclude              bool            `json:"exclude"`
	Period               *Period         `json:"period,omitempty"`
}

// Identifies the resource instances that are members of the group.
type GroupMember struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Entity            Reference   `json:"entity"`
	Period            *Period     `json:"period,omitempty"`
	Inactive          *bool       `json:"inactive,omitempty"`
}

// This function returns resource reference information
func (r Group) ResourceRef() (string, *string) {
	return "Group", r.ID
}

// This function returns resource reference information
func (r Group) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherGroup Group

// MarshalJSON marshals the given Group as JSON into a byte slice
func (r Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherGroup:   OtherGroup(r),
		ResourceType: "Group",
	})
}

// UnmarshalGroup unmarshals a Group.
func UnmarshalGroup(b []byte) (Group, error) {
	var group Group
	if err := json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}
