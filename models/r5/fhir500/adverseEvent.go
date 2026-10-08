package fhir500

import (
	"encoding/json"

	"github.com/savannahghi/scalarutils"
)

// AdverseEvent is documented here https://hl7.org/fhir/R5/adverseevent.html
// An event (i.e. any change to current patient status) that may be related to
// unintended effects on a patient or research participant. The unintended
// effects may require additional monitoring, treatment, hospitalization, or may
// result in death.
type AdverseEvent struct {
	ID                      *string                           `json:"id,omitempty"`
	Meta                    *Meta                             `json:"meta,omitempty"`
	ImplicitRules           *string                           `json:"implicitRules,omitempty"`
	Language                *string                           `json:"language,omitempty"`
	Text                    *Narrative                        `json:"text,omitempty"`
	Extension               []Extension                       `json:"extension,omitempty"`
	ModifierExtension       []Extension                       `json:"modifierExtension,omitempty"`
	Identifier              []*Identifier                     `json:"identifier,omitempty"`
	Status                  *AdverseEventStatus               `json:"status,omitempty"`
	Actuality               *AdverseEventActuality            `json:"actuality,omitempty"`
	Category                []*CodeableConcept                `json:"category,omitempty"`
	Code                    *CodeableConcept                  `json:"code,omitempty"`
	Subject                 *Reference                        `json:"subject,omitempty"`
	Encounter               *Reference                        `json:"encounter,omitempty"`
	OccurrenceDateTime      *scalarutils.Date                 `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period                           `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                           `json:"occurrenceTiming,omitempty"`
	Detected                *scalarutils.Date                 `json:"detected,omitempty"`
	RecordedDate            *scalarutils.Date                 `json:"recordedDate,omitempty"`
	ResultingEffect         []*Reference                      `json:"resultingEffect,omitempty"`
	Location                *Reference                        `json:"location,omitempty"`
	Seriousness             *CodeableConcept                  `json:"seriousness,omitempty"`
	Outcome                 []*CodeableConcept                `json:"outcome,omitempty"`
	Recorder                *Reference                        `json:"recorder,omitempty"`
	Participant             []*AdverseEventParticipant        `json:"participant,omitempty"`
	Study                   []*Reference                      `json:"study,omitempty"`
	ExpectedInResearchStudy *bool                             `json:"expectedInResearchStudy,omitempty"`
	SuspectEntity           []*AdverseEventSuspectEntity      `json:"suspectEntity,omitempty"`
	ContributingFactor      []*AdverseEventContributingFactor `json:"contributingFactor,omitempty"`
	PreventiveAction        []*AdverseEventPreventiveAction   `json:"preventiveAction,omitempty"`
	MitigatingAction        []*AdverseEventMitigatingAction   `json:"mitigatingAction,omitempty"`
	SupportingInfo          []*AdverseEventSupportingInfo     `json:"supportingInfo,omitempty"`
	Note                    []*Annotation                     `json:"note,omitempty"`
}

// AdverseEventParticipant indicates who or what participated in the adverse
// event and how they were involved.
//
// R5 replaced the R4B contributor reference list with this backbone element so a
// participant's role can be qualified by function.
type AdverseEventParticipant struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             *Reference       `json:"actor,omitempty"`
}

// AdverseEventSuspectEntity describes the entity that is suspected to have
// caused the adverse event.
type AdverseEventSuspectEntity struct {
	ID                      *string                             `json:"id,omitempty"`
	Extension               []Extension                         `json:"extension,omitempty"`
	ModifierExtension       []Extension                         `json:"modifierExtension,omitempty"`
	InstanceCodeableConcept *CodeableConcept                    `json:"instanceCodeableConcept,omitempty"`
	InstanceReference       *Reference                          `json:"instanceReference,omitempty"`
	Causality               *AdverseEventSuspectEntityCausality `json:"causality,omitempty"`
}

// AdverseEventSuspectEntityCausality holds information on the possible cause of
// the event.
//
// R5 replaced the R4B assessment, productRelatedness and method elements with
// assessmentMethod and entityRelatedness.
type AdverseEventSuspectEntityCausality struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	AssessmentMethod  *CodeableConcept `json:"assessmentMethod,omitempty"`
	EntityRelatedness *CodeableConcept `json:"entityRelatedness,omitempty"`
	Author            *Reference       `json:"author,omitempty"`
}

// AdverseEventContributingFactor records the contributing factors suspected to
// have increased the probability or severity of the adverse event.
type AdverseEventContributingFactor struct {
	ID                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventPreventiveAction records the preventive actions that contributed
// to avoiding the adverse event.
type AdverseEventPreventiveAction struct {
	ID                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventMitigatingAction records the ameliorating actions taken after the
// adverse event to reduce the extent of harm.
type AdverseEventMitigatingAction struct {
	ID                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventSupportingInfo holds supporting information relevant to the event.
//
// R5 folded the R4B subjectMedicalHistory and referenceDocument lists into this
// single element.
type AdverseEventSupportingInfo struct {
	ID                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventRelayPayload is the wrapper used when an AdverseEvent is relayed
// inside another payload.
type AdverseEventRelayPayload struct {
	Resource *AdverseEvent `json:"resource,omitempty"`
}

type OtherAdverseEvent AdverseEvent

// MarshalJSON marshals the given AdverseEvent as JSON into a byte slice.
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}

// UnmarshalAdverseEvent unmarshals an AdverseEvent.
func UnmarshalAdverseEvent(b []byte) (AdverseEvent, error) {
	var adverseEvent AdverseEvent
	if err := json.Unmarshal(b, &adverseEvent); err != nil {
		return adverseEvent, err
	}

	return adverseEvent, nil
}
