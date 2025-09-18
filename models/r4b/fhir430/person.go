
package fhir430

import "encoding/json"
// Person is documented here http://hl7.org/fhir/StructureDefinition/Person
// Demographics and administrative information about a person independent of a specific health-related context.
type Person struct {
	ID                   *string               `json:"ID,omitempty"`
	Meta                 *Meta                 `json:"meta,omitempty"`
	ImplicitRules        *string               `json:"implicitRules,omitempty"`
	Language             *string               `json:"language,omitempty"`
	Text                 *Narrative            `json:"text,omitempty"`
	Contained            []json.RawMessage     `json:"contained,omitempty"`
	Extension            []Extension           `json:"extension,omitempty"`
	ModifierExtension    []Extension           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `json:"identifier,omitempty"`
	Name                 []HumanName           `json:"name,omitempty"`
	Telecom              []ContactPoint        `json:"telecom,omitempty"`
	Gender               *AdministrativeGender `json:"gender,omitempty"`
	BirthDate            *string               `json:"birthDate,omitempty"`
	Address              []Address             `json:"address,omitempty"`
	Photo                *Attachment           `json:"photo,omitempty"`
	ManagingOrganization *Reference            `json:"managingOrganization,omitempty"`
	Active               *bool                 `json:"active,omitempty"`
	Link                 []PersonLink          `json:"link,omitempty"`
}

// Link to a resource that concerns the same actual person.
type PersonLink struct {
	ID                *string                 `json:"ID,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Target            Reference               `json:"target"`
	Assurance         *IdentityAssuranceLevel `json:"assurance,omitempty"`
}

// This function returns resource reference information
func (r Person) ResourceRef() (string, *string) {
	return "Person", r.ID
}

// This function returns resource reference information
func (r Person) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherPerson Person

// MarshalJSON marshals the given Person as JSON into a byte slice
func (r Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherPerson:  OtherPerson(r),
		ResourceType: "Person",
	})
}

// UnmarshalPerson unmarshals a Person.
func UnmarshalPerson(b []byte) (Person, error) {
	var person Person
	if err := json.Unmarshal(b, &person); err != nil {
		return person, err
	}
	return person, nil
}
