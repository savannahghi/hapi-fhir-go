package fhir430

import "encoding/json"

// Media is documented here http://hl7.org/fhir/StructureDefinition/Media
// A photo, video, or audio recording acquired or used in healthcare. The actual content may be inline or provided by direct reference.
type Media struct {
	ID                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	BasedOn           []Reference       `json:"basedOn,omitempty"`
	PartOf            []Reference       `json:"partOf,omitempty"`
	Status            EventStatus       `json:"status"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Modality          *CodeableConcept  `json:"modality,omitempty"`
	View              *CodeableConcept  `json:"view,omitempty"`
	Subject           *Reference        `json:"subject,omitempty"`
	Encounter         *Reference        `json:"encounter,omitempty"`
	CreatedDateTime   *string           `json:"createdDateTime,omitempty"`
	CreatedPeriod     *Period           `json:"createdPeriod,omitempty"`
	Issued            *string           `json:"issued,omitempty"`
	Operator          *Reference        `json:"operator,omitempty"`
	ReasonCode        []CodeableConcept `json:"reasonCode,omitempty"`
	BodySite          *CodeableConcept  `json:"bodySite,omitempty"`
	DeviceName        *string           `json:"deviceName,omitempty"`
	Device            *Reference        `json:"device,omitempty"`
	Height            *int              `json:"height,omitempty"`
	Width             *int              `json:"width,omitempty"`
	Frames            *int              `json:"frames,omitempty"`
	Duration          *json.Number      `json:"duration,omitempty"`
	Content           Attachment        `json:"content"`
	Note              []Annotation      `json:"note,omitempty"`
}

// This function returns resource reference information
func (r Media) ResourceRef() (string, *string) {
	return "Media", r.ID
}

// This function returns resource reference information
func (r Media) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedia Media

// MarshalJSON marshals the given Media as JSON into a byte slice
func (r Media) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedia
		ResourceType string `json:"resourceType"`
	}{
		OtherMedia:   OtherMedia(r),
		ResourceType: "Media",
	})
}

// UnmarshalMedia unmarshals a Media.
func UnmarshalMedia(b []byte) (Media, error) {
	var media Media
	if err := json.Unmarshal(b, &media); err != nil {
		return media, err
	}
	return media, nil
}
