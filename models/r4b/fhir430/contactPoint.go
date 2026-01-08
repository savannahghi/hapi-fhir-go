package fhir430

// ContactPoint is documented here http://hl7.org/fhir/StructureDefinition/ContactPoint
// Base StructureDefinition for ContactPoint Type: Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
type ContactPoint struct {
	ID        *string             `json:"id,omitempty"`
	Extension []Extension         `json:"extension,omitempty"`
	System    *ContactPointSystem `json:"system,omitempty"`
	Value     *string             `json:"value,omitempty"`
	Use       *ContactPointUse    `json:"use,omitempty"`
	Rank      *int                `json:"rank,omitempty"`
	Period    *Period             `json:"period,omitempty"`
}
