package fhir430

import "encoding/json"

// DocumentManifest is documented here http://hl7.org/fhir/StructureDefinition/DocumentManifest
// A collection of documents compiled for a purpose together with metadata that applies to the collection.
type DocumentManifest struct {
	ID                *string                   `json:"id,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []json.RawMessage         `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier               `json:"masterIdentifier,omitempty"`
	Identifier        []Identifier              `json:"identifier,omitempty"`
	Status            DocumentReferenceStatus   `json:"status"`
	Type              *CodeableConcept          `json:"type,omitempty"`
	Subject           *Reference                `json:"subject,omitempty"`
	Created           *string                   `json:"created,omitempty"`
	Author            []Reference               `json:"author,omitempty"`
	Recipient         []Reference               `json:"recipient,omitempty"`
	Source            *string                   `json:"source,omitempty"`
	Description       *string                   `json:"description,omitempty"`
	Content           []Reference               `json:"content"`
	Related           []DocumentManifestRelated `json:"related,omitempty"`
}

// Related identifiers or resources associated with the DocumentManifest.
// May be identifiers or resources that caused the DocumentManifest to be created.
type DocumentManifestRelated struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Ref               *Reference  `json:"ref,omitempty"`
}

// This function returns resource reference information
func (r DocumentManifest) ResourceRef() (string, *string) {
	return "DocumentManifest", r.ID
}

// This function returns resource reference information
func (r DocumentManifest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDocumentManifest DocumentManifest

// MarshalJSON marshals the given DocumentManifest as JSON into a byte slice
func (r DocumentManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentManifest
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentManifest: OtherDocumentManifest(r),
		ResourceType:          "DocumentManifest",
	})
}

// UnmarshalDocumentManifest unmarshals a DocumentManifest.
func UnmarshalDocumentManifest(b []byte) (DocumentManifest, error) {
	var documentManifest DocumentManifest
	if err := json.Unmarshal(b, &documentManifest); err != nil {
		return documentManifest, err
	}
	return documentManifest, nil
}
