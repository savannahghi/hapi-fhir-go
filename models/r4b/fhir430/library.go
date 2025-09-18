
package fhir430

import "encoding/json"
// Library is documented here http://hl7.org/fhir/StructureDefinition/Library
// The Library resource is a general-purpose container for knowledge asset definitions. It can be used to describe and expose existing knowledge assets such as logic libraries and information model descriptions, as well as to describe a collection of knowledge assets.
type Library struct {
	ID                     *string               `json:"ID,omitempty"`
	Meta                   *Meta                 `json:"meta,omitempty"`
	ImplicitRules          *string               `json:"implicitRules,omitempty"`
	Language               *string               `json:"language,omitempty"`
	Text                   *Narrative            `json:"text,omitempty"`
	Contained              []json.RawMessage     `json:"contained,omitempty"`
	Extension              []Extension           `json:"extension,omitempty"`
	ModifierExtension      []Extension           `json:"modifierExtension,omitempty"`
	Url                    *string               `json:"url,omitempty"`
	Identifier             []Identifier          `json:"identifier,omitempty"`
	Version                *string               `json:"version,omitempty"`
	Name                   *string               `json:"name,omitempty"`
	Title                  *string               `json:"title,omitempty"`
	Subtitle               *string               `json:"subtitle,omitempty"`
	Status                 PublicationStatus     `json:"status"`
	Experimental           *bool                 `json:"experimental,omitempty"`
	Type                   CodeableConcept       `json:"type"`
	SubjectCodeableConcept *CodeableConcept      `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference            `json:"subjectReference,omitempty"`
	Date                   *string               `json:"date,omitempty"`
	Publisher              *string               `json:"publisher,omitempty"`
	Contact                []ContactDetail       `json:"contact,omitempty"`
	Description            *string               `json:"description,omitempty"`
	UseContext             []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose                *string               `json:"purpose,omitempty"`
	Usage                  *string               `json:"usage,omitempty"`
	ApprovalDate           *string               `json:"approvalDate,omitempty"`
	LastReviewDate         *string               `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period               `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept     `json:"topic,omitempty"`
	Author                 []ContactDetail       `json:"author,omitempty"`
	Editor                 []ContactDetail       `json:"editor,omitempty"`
	Reviewer               []ContactDetail       `json:"reviewer,omitempty"`
	Endorser               []ContactDetail       `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact     `json:"relatedArtifact,omitempty"`
	Parameter              []ParameterDefinition `json:"parameter,omitempty"`
	DataRequirement        []DataRequirement     `json:"dataRequirement,omitempty"`
	Content                []Attachment          `json:"content,omitempty"`
}

// This function returns resource reference information
func (r Library) ResourceRef() (string, *string) {
	return "Library", r.ID
}

// This function returns resource reference information
func (r Library) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherLibrary Library

// MarshalJSON marshals the given Library as JSON into a byte slice
func (r Library) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLibrary
		ResourceType string `json:"resourceType"`
	}{
		OtherLibrary: OtherLibrary(r),
		ResourceType: "Library",
	})
}

// UnmarshalLibrary unmarshals a Library.
func UnmarshalLibrary(b []byte) (Library, error) {
	var library Library
	if err := json.Unmarshal(b, &library); err != nil {
		return library, err
	}
	return library, nil
}
