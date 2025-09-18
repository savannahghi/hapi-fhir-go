
package fhir430
// Annotation is documented here http://hl7.org/fhir/StructureDefinition/Annotation
// Base StructureDefinition for Annotation Type: A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	ID              *string     `json:"ID,omitempty"`
	Extension       []Extension `json:"extension,omitempty"`
	AuthorReference *Reference  `json:"authorReference,omitempty"`
	AuthorString    *string     `json:"authorString,omitempty"`
	Time            *string     `json:"time,omitempty"`
	Text            string      `json:"text"`
}
