
package fhir430

import "encoding/json"
// List is documented here http://hl7.org/fhir/StructureDefinition/List
// A list is a curated collection of resources.
type List struct {
	ID                *string           `json:"ID,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Status            ListStatus        `json:"status"`
	Mode              ListMode          `json:"mode"`
	Title             *string           `json:"title,omitempty"`
	Code              *CodeableConcept  `json:"code,omitempty"`
	Subject           *Reference        `json:"subject,omitempty"`
	Encounter         *Reference        `json:"encounter,omitempty"`
	Date              *string           `json:"date,omitempty"`
	Source            *Reference        `json:"source,omitempty"`
	OrderedBy         *CodeableConcept  `json:"orderedBy,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
	Entry             []ListEntry       `json:"entry,omitempty"`
	EmptyReason       *CodeableConcept  `json:"emptyReason,omitempty"`
}

// Entries in this list.
// If there are no entries in the list, an emptyReason SHOULD be provided.
type ListEntry struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Flag              *CodeableConcept `json:"flag,omitempty"`
	Deleted           *bool            `json:"deleted,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Item              Reference        `json:"item"`
}

// This function returns resource reference information
func (r List) ResourceRef() (string, *string) {
	return "List", r.ID
}

// This function returns resource reference information
func (r List) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherList List

// MarshalJSON marshals the given List as JSON into a byte slice
func (r List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherList
		ResourceType string `json:"resourceType"`
	}{
		OtherList:    OtherList(r),
		ResourceType: "List",
	})
}

// UnmarshalList unmarshals a List.
func UnmarshalList(b []byte) (List, error) {
	var list List
	if err := json.Unmarshal(b, &list); err != nil {
		return list, err
	}
	return list, nil
}
