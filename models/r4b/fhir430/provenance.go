
package fhir430

import "encoding/json"
// Provenance is documented here http://hl7.org/fhir/StructureDefinition/Provenance
// Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies.
type Provenance struct {
	ID                *string            `json:"ID,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []json.RawMessage  `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Target            []Reference        `json:"target"`
	OccurredPeriod    *Period            `json:"occurredPeriod,omitempty"`
	OccurredDateTime  *string            `json:"occurredDateTime,omitempty"`
	Recorded          string             `json:"recorded"`
	Policy            []string           `json:"policy,omitempty"`
	Location          *Reference         `json:"location,omitempty"`
	Reason            []CodeableConcept  `json:"reason,omitempty"`
	Activity          *CodeableConcept   `json:"activity,omitempty"`
	Agent             []ProvenanceAgent  `json:"agent"`
	Entity            []ProvenanceEntity `json:"entity,omitempty"`
	Signature         []Signature        `json:"signature,omitempty"`
}

// An actor taking a role in an activity  for which it can be assigned some degree of responsibility for the activity taking place.
// Several agents may be associated (i.e. has some responsibility for an activity) with an activity and vice-versa.
type ProvenanceAgent struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Who               Reference         `json:"who"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
}

// An entity used in this activity.
type ProvenanceEntity struct {
	ID                *string              `json:"ID,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Role              ProvenanceEntityRole `json:"role"`
	What              Reference            `json:"what"`
	Agent             []ProvenanceAgent    `json:"agent,omitempty"`
}

// This function returns resource reference information
func (r Provenance) ResourceRef() (string, *string) {
	return "Provenance", r.ID
}

// This function returns resource reference information
func (r Provenance) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherProvenance Provenance

// MarshalJSON marshals the given Provenance as JSON into a byte slice
func (r Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProvenance
		ResourceType string `json:"resourceType"`
	}{
		OtherProvenance: OtherProvenance(r),
		ResourceType:    "Provenance",
	})
}

// UnmarshalProvenance unmarshals a Provenance.
func UnmarshalProvenance(b []byte) (Provenance, error) {
	var provenance Provenance
	if err := json.Unmarshal(b, &provenance); err != nil {
		return provenance, err
	}
	return provenance, nil
}
