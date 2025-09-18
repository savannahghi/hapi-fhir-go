
package fhir430

import "encoding/json"
// Linkage is documented here http://hl7.org/fhir/StructureDefinition/Linkage
// Identifies two or more records (resource instances) that refer to the same real-world "occurrence".
type Linkage struct {
	ID                *string           `json:"ID,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Active            *bool             `json:"active,omitempty"`
	Author            *Reference        `json:"author,omitempty"`
	Item              []LinkageItem     `json:"item"`
}

// Identifies which record considered as the reference to the same real-world occurrence as well as how the items should be evaluated within the collection of linked items.
type LinkageItem struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              LinkageType `json:"type"`
	Resource          Reference   `json:"resource"`
}

// This function returns resource reference information
func (r Linkage) ResourceRef() (string, *string) {
	return "Linkage", r.ID
}

// This function returns resource reference information
func (r Linkage) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherLinkage Linkage

// MarshalJSON marshals the given Linkage as JSON into a byte slice
func (r Linkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLinkage
		ResourceType string `json:"resourceType"`
	}{
		OtherLinkage: OtherLinkage(r),
		ResourceType: "Linkage",
	})
}

// UnmarshalLinkage unmarshals a Linkage.
func UnmarshalLinkage(b []byte) (Linkage, error) {
	var linkage Linkage
	if err := json.Unmarshal(b, &linkage); err != nil {
		return linkage, err
	}
	return linkage, nil
}
