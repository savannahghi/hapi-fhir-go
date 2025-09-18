
package fhir430

import "encoding/json"
// Composition is documented here http://hl7.org/fhir/StructureDefinition/Composition
// A set of healthcare-related information that is assembled together into a single logical package that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. A Composition defines the structure and narrative content necessary for a document. However, a Composition alone does not constitute a document. Rather, the Composition must be the first entry in a Bundle where Bundle.type=document, and any other resources referenced from Composition must be included as subsequent entries in the Bundle (for example Patient, Practitioner, Encounter, etc.).
type Composition struct {
	ID                *string                `json:"ID,omitempty"`
	Meta              *Meta                  `json:"meta,omitempty"`
	ImplicitRules     *string                `json:"implicitRules,omitempty"`
	Language          *string                `json:"language,omitempty"`
	Text              *Narrative             `json:"text,omitempty"`
	Contained         []json.RawMessage      `json:"contained,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Identifier        *Identifier            `json:"identifier,omitempty"`
	Status            CompositionStatus      `json:"status"`
	Type              CodeableConcept        `json:"type"`
	Category          []CodeableConcept      `json:"category,omitempty"`
	Subject           *Reference             `json:"subject,omitempty"`
	Encounter         *Reference             `json:"encounter,omitempty"`
	Date              string                 `json:"date"`
	Author            []Reference            `json:"author"`
	Title             string                 `json:"title"`
	Confidentiality   *string                `json:"confidentiality,omitempty"`
	Attester          []CompositionAttester  `json:"attester,omitempty"`
	Custodian         *Reference             `json:"custodian,omitempty"`
	RelatesTo         []CompositionRelatesTo `json:"relatesTo,omitempty"`
	Event             []CompositionEvent     `json:"event,omitempty"`
	Section           []CompositionSection   `json:"section,omitempty"`
}

// A participant who has attested to the accuracy of the composition/document.
// Only list each attester once.
type CompositionAttester struct {
	ID                *string                    `json:"ID,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Mode              CompositionAttestationMode `json:"mode"`
	Time              *string                    `json:"time,omitempty"`
	Party             *Reference                 `json:"party,omitempty"`
}

// Relationships that this composition has with other compositions or documents that already exist.
// A document is a version specific composition.
type CompositionRelatesTo struct {
	ID                *string                  `json:"ID,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Code              DocumentRelationshipType `json:"code"`
	TargetIdentifier  Identifier               `json:"targetIdentifier"`
	TargetReference   Reference                `json:"targetReference"`
}

// The clinical service, such as a colonoscopy or an appendectomy, being documented.
// The event needs to be consistent with the type element, though can provide further information if desired.
type CompositionEvent struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

// The root of the sections that make up the composition.
type CompositionSection struct {
	ID                *string              `json:"ID,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Title             *string              `json:"title,omitempty"`
	Code              *CodeableConcept     `json:"code,omitempty"`
	Author            []Reference          `json:"author,omitempty"`
	Focus             *Reference           `json:"focus,omitempty"`
	Text              *Narrative           `json:"text,omitempty"`
	Mode              *ListMode            `json:"mode,omitempty"`
	OrderedBy         *CodeableConcept     `json:"orderedBy,omitempty"`
	Entry             []Reference          `json:"entry,omitempty"`
	EmptyReason       *CodeableConcept     `json:"emptyReason,omitempty"`
	Section           []CompositionSection `json:"section,omitempty"`
}

// This function returns resource reference information
func (r Composition) ResourceRef() (string, *string) {
	return "Composition", r.ID
}

// This function returns resource reference information
func (r Composition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherComposition Composition

// MarshalJSON marshals the given Composition as JSON into a byte slice
func (r Composition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherComposition
		ResourceType string `json:"resourceType"`
	}{
		OtherComposition: OtherComposition(r),
		ResourceType:     "Composition",
	})
}

// UnmarshalComposition unmarshals a Composition.
func UnmarshalComposition(b []byte) (Composition, error) {
	var composition Composition
	if err := json.Unmarshal(b, &composition); err != nil {
		return composition, err
	}
	return composition, nil
}
