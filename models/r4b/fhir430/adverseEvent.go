package fhir430

import "encoding/json"

// AdverseEvent is documented here http://hl7.org/fhir/StructureDefinition/AdverseEvent
// Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death.
type AdverseEvent struct {
	ID                    *string                     `json:"id,omitempty"`
	Meta                  *Meta                       `json:"meta,omitempty"`
	ImplicitRules         *string                     `json:"implicitRules,omitempty"`
	Language              *string                     `json:"language,omitempty"`
	Text                  *Narrative                  `json:"text,omitempty"`
	Contained             []json.RawMessage           `json:"contained,omitempty"`
	Extension             []Extension                 `json:"extension,omitempty"`
	ModifierExtension     []Extension                 `json:"modifierExtension,omitempty"`
	Identifier            *Identifier                 `json:"identifier,omitempty"`
	Actuality             AdverseEventActuality       `json:"actuality"`
	Category              []CodeableConcept           `json:"category,omitempty"`
	Event                 *CodeableConcept            `json:"event,omitempty"`
	Subject               Reference                   `json:"subject"`
	Encounter             *Reference                  `json:"encounter,omitempty"`
	Date                  *string                     `json:"date,omitempty"`
	Detected              *string                     `json:"detected,omitempty"`
	RecordedDate          *string                     `json:"recordedDate,omitempty"`
	ResultingCondition    []Reference                 `json:"resultingCondition,omitempty"`
	Location              *Reference                  `json:"location,omitempty"`
	Seriousness           *CodeableConcept            `json:"seriousness,omitempty"`
	Severity              *CodeableConcept            `json:"severity,omitempty"`
	Outcome               *CodeableConcept            `json:"outcome,omitempty"`
	Recorder              *Reference                  `json:"recorder,omitempty"`
	Contributor           []Reference                 `json:"contributor,omitempty"`
	SuspectEntity         []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
	SubjectMedicalHistory []Reference                 `json:"subjectMedicalHistory,omitempty"`
	ReferenceDocument     []Reference                 `json:"referenceDocument,omitempty"`
	Study                 []Reference                 `json:"study,omitempty"`
}

// Describes the entity that is suspected to have caused the adverse event.
type AdverseEventSuspectEntity struct {
	ID                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Instance          Reference                            `json:"instance"`
	Causality         []AdverseEventSuspectEntityCausality `json:"causality,omitempty"`
}

// Information on the possible cause of the event.
type AdverseEventSuspectEntityCausality struct {
	ID                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Assessment         *CodeableConcept `json:"assessment,omitempty"`
	ProductRelatedness *string          `json:"productRelatedness,omitempty"`
	Author             *Reference       `json:"author,omitempty"`
	Method             *CodeableConcept `json:"method,omitempty"`
}

// This function returns resource reference information
func (r AdverseEvent) ResourceRef() (string, *string) {
	return "AdverseEvent", r.ID
}

// This function returns resource reference information
func (r AdverseEvent) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherAdverseEvent AdverseEvent

// MarshalJSON marshals the given AdverseEvent as JSON into a byte slice
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}

// UnmarshalAdverseEvent unmarshals a AdverseEvent.
func UnmarshalAdverseEvent(b []byte) (AdverseEvent, error) {
	var adverseEvent AdverseEvent
	if err := json.Unmarshal(b, &adverseEvent); err != nil {
		return adverseEvent, err
	}
	return adverseEvent, nil
}
