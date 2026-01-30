package fhir430

import "encoding/json"

// Procedure is documented here http://hl7.org/fhir/StructureDefinition/Procedure
// An action that is or was performed on or for a patient. This can be a physical intervention like an operation, or less invasive like long term services, counseling, or hypnotherapy.
type Procedure struct {
	ID                    *string                `json:"id,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Contained             []json.RawMessage      `json:"contained,omitempty"`
	Extension             []Extension            `json:"extension,omitempty"`
	ModifierExtension     []Extension            `json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical []string               `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string               `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference            `json:"basedOn,omitempty"`
	PartOf                []Reference            `json:"partOf,omitempty"`
	Status                EventStatus            `json:"status"`
	StatusReason          *CodeableConcept       `json:"statusReason,omitempty"`
	Category              *CodeableConcept       `json:"category,omitempty"`
	Code                  *CodeableConcept       `json:"code,omitempty"`
	Subject               Reference              `json:"subject"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	PerformedDateTime     *string                `json:"performedDateTime,omitempty"`
	PerformedPeriod       *Period                `json:"performedPeriod,omitempty"`
	PerformedString       *string                `json:"performedString,omitempty"`
	PerformedAge          *Age                   `json:"performedAge,omitempty"`
	PerformedRange        *Range                 `json:"performedRange,omitempty"`
	Recorder              *Reference             `json:"recorder,omitempty"`
	Asserter              *Reference             `json:"asserter,omitempty"`
	Performer             []ProcedurePerformer   `json:"performer,omitempty"`
	Location              *Reference             `json:"location,omitempty"`
	ReasonCode            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `json:"reasonReference,omitempty"`
	BodySite              []CodeableConcept      `json:"bodySite,omitempty"`
	Outcome               *CodeableConcept       `json:"outcome,omitempty"`
	Report                []Reference            `json:"report,omitempty"`
	Complication          []CodeableConcept      `json:"complication,omitempty"`
	ComplicationDetail    []Reference            `json:"complicationDetail,omitempty"`
	FollowUp              []CodeableConcept      `json:"followUp,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
	FocalDevice           []ProcedureFocalDevice `json:"focalDevice,omitempty"`
	UsedReference         []Reference            `json:"usedReference,omitempty"`
	UsedCode              []CodeableConcept      `json:"usedCode,omitempty"`
}

// Limited to "real" people rather than equipment.
type ProcedurePerformer struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
	OnBehalfOf        *Reference       `json:"onBehalfOf,omitempty"`
}

// A device that is implanted, removed or otherwise manipulated (calibration, battery replacement, fitting a prosthesis, attaching a wound-vac, etc.) as a focal portion of the Procedure.
type ProcedureFocalDevice struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Action            *CodeableConcept `json:"action,omitempty"`
	Manipulated       Reference        `json:"manipulated"`
}

// This function returns resource reference information
func (r Procedure) ResourceRef() (string, *string) {
	return "Procedure", r.ID
}

// This function returns resource reference information
func (r Procedure) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherProcedure Procedure

// MarshalJSON marshals the given Procedure as JSON into a byte slice
func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcedure
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
}

// UnmarshalProcedure unmarshals a Procedure.
func UnmarshalProcedure(b []byte) (Procedure, error) {
	var procedure Procedure
	if err := json.Unmarshal(b, &procedure); err != nil {
		return procedure, err
	}
	return procedure, nil
}
