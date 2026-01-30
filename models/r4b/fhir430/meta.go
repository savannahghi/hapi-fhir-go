package fhir430

// Meta is documented here http://hl7.org/fhir/StructureDefinition/Meta
// Base StructureDefinition for Meta Type: The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
type Meta struct {
	ID          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	VersionId   *string     `json:"versionId,omitempty"`
	LastUpdated *string     `json:"lastUpdated,omitempty"`
	Source      *string     `json:"source,omitempty"`
	Profile     []string    `json:"profile,omitempty"`
	Security    []Coding    `json:"security,omitempty"`
	Tag         []Coding    `json:"tag,omitempty"`
}
