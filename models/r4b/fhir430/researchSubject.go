package fhir430

import "encoding/json"

// ResearchSubject is documented here http://hl7.org/fhir/StructureDefinition/ResearchSubject
// A physical entity which is the primary unit of operational and/or administrative interest in a study.
type ResearchSubject struct {
	ID                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []json.RawMessage     `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Status            ResearchSubjectStatus `json:"status"`
	Period            *Period               `json:"period,omitempty"`
	Study             Reference             `json:"study"`
	Individual        Reference             `json:"individual"`
	AssignedArm       *string               `json:"assignedArm,omitempty"`
	ActualArm         *string               `json:"actualArm,omitempty"`
	Consent           *Reference            `json:"consent,omitempty"`
}

// This function returns resource reference information
func (r ResearchSubject) ResourceRef() (string, *string) {
	return "ResearchSubject", r.ID
}

// This function returns resource reference information
func (r ResearchSubject) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherResearchSubject ResearchSubject

// MarshalJSON marshals the given ResearchSubject as JSON into a byte slice
func (r ResearchSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchSubject
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchSubject: OtherResearchSubject(r),
		ResourceType:         "ResearchSubject",
	})
}

// UnmarshalResearchSubject unmarshals a ResearchSubject.
func UnmarshalResearchSubject(b []byte) (ResearchSubject, error) {
	var researchSubject ResearchSubject
	if err := json.Unmarshal(b, &researchSubject); err != nil {
		return researchSubject, err
	}
	return researchSubject, nil
}
