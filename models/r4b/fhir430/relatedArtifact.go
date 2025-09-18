
package fhir430
// RelatedArtifact is documented here http://hl7.org/fhir/StructureDefinition/RelatedArtifact
// Base StructureDefinition for RelatedArtifact Type: Related artifacts such as additional documentation, justification, or bibliographic references.
type RelatedArtifact struct {
	ID        *string             `json:"ID,omitempty"`
	Extension []Extension         `json:"extension,omitempty"`
	Type      RelatedArtifactType `json:"type"`
	Label     *string             `json:"label,omitempty"`
	Display   *string             `json:"display,omitempty"`
	Citation  *string             `json:"citation,omitempty"`
	Url       *string             `json:"url,omitempty"`
	Document  *Attachment         `json:"document,omitempty"`
	Resource  *string             `json:"resource,omitempty"`
}
