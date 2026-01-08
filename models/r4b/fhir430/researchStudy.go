package fhir430

import "encoding/json"

// ResearchStudy is documented here http://hl7.org/fhir/StructureDefinition/ResearchStudy
// A process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge.  This includes studies of safety, efficacy, comparative effectiveness and other information about medications, devices, therapies and other interventional and investigative techniques.  A ResearchStudy involves the gathering of information about human or animal subjects.
type ResearchStudy struct {
	ID                    *string                  `json:"id,omitempty"`
	Meta                  *Meta                    `json:"meta,omitempty"`
	ImplicitRules         *string                  `json:"implicitRules,omitempty"`
	Language              *string                  `json:"language,omitempty"`
	Text                  *Narrative               `json:"text,omitempty"`
	Contained             []json.RawMessage        `json:"contained,omitempty"`
	Extension             []Extension              `json:"extension,omitempty"`
	ModifierExtension     []Extension              `json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `json:"identifier,omitempty"`
	Title                 *string                  `json:"title,omitempty"`
	Protocol              []Reference              `json:"protocol,omitempty"`
	PartOf                []Reference              `json:"partOf,omitempty"`
	Status                ResearchStudyStatus      `json:"status"`
	PrimaryPurposeType    *CodeableConcept         `json:"primaryPurposeType,omitempty"`
	Phase                 *CodeableConcept         `json:"phase,omitempty"`
	Category              []CodeableConcept        `json:"category,omitempty"`
	Focus                 []CodeableConcept        `json:"focus,omitempty"`
	Condition             []CodeableConcept        `json:"condition,omitempty"`
	Contact               []ContactDetail          `json:"contact,omitempty"`
	RelatedArtifact       []RelatedArtifact        `json:"relatedArtifact,omitempty"`
	Keyword               []CodeableConcept        `json:"keyword,omitempty"`
	Location              []CodeableConcept        `json:"location,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	Enrollment            []Reference              `json:"enrollment,omitempty"`
	Period                *Period                  `json:"period,omitempty"`
	Sponsor               *Reference               `json:"sponsor,omitempty"`
	PrincipalInvestigator *Reference               `json:"principalInvestigator,omitempty"`
	Site                  []Reference              `json:"site,omitempty"`
	ReasonStopped         *CodeableConcept         `json:"reasonStopped,omitempty"`
	Note                  []Annotation             `json:"note,omitempty"`
	Arm                   []ResearchStudyArm       `json:"arm,omitempty"`
	Objective             []ResearchStudyObjective `json:"objective,omitempty"`
}

// Describes an expected sequence of events for one of the participants of a study.  E.g. Exposure to drug A, wash-out, exposure to drug B, wash-out, follow-up.
type ResearchStudyArm struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              string           `json:"name"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Description       *string          `json:"description,omitempty"`
}

// A goal that the study is aiming to achieve in terms of a scientific question to be answered by the analysis of data collected during the study.
type ResearchStudyObjective struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

// This function returns resource reference information
func (r ResearchStudy) ResourceRef() (string, *string) {
	return "ResearchStudy", r.ID
}

// This function returns resource reference information
func (r ResearchStudy) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherResearchStudy ResearchStudy

// MarshalJSON marshals the given ResearchStudy as JSON into a byte slice
func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchStudy: OtherResearchStudy(r),
		ResourceType:       "ResearchStudy",
	})
}

// UnmarshalResearchStudy unmarshals a ResearchStudy.
func UnmarshalResearchStudy(b []byte) (ResearchStudy, error) {
	var researchStudy ResearchStudy
	if err := json.Unmarshal(b, &researchStudy); err != nil {
		return researchStudy, err
	}
	return researchStudy, nil
}
