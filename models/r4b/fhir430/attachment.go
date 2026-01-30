package fhir430

// Attachment is documented here http://hl7.org/fhir/StructureDefinition/Attachment
// Base StructureDefinition for Attachment Type: For referring to data content defined in other formats.
type Attachment struct {
	ID          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	ContentType *string     `json:"contentType,omitempty"`
	Language    *string     `json:"language,omitempty"`
	Data        *string     `json:"data,omitempty"`
	Url         *string     `json:"url,omitempty"`
	Size        *int        `json:"size,omitempty"`
	Hash        *string     `json:"hash,omitempty"`
	Title       *string     `json:"title,omitempty"`
	Creation    *string     `json:"creation,omitempty"`
}
