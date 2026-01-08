package fhir430

import "encoding/json"

// ActivityDefinition is documented here http://hl7.org/fhir/StructureDefinition/ActivityDefinition
// This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinition struct {
	ID                           *string                          `json:"id,omitempty"`
	Meta                         *Meta                            `json:"meta,omitempty"`
	ImplicitRules                *string                          `json:"implicitRules,omitempty"`
	Language                     *string                          `json:"language,omitempty"`
	Text                         *Narrative                       `json:"text,omitempty"`
	Contained                    []json.RawMessage                `json:"contained,omitempty"`
	Extension                    []Extension                      `json:"extension,omitempty"`
	ModifierExtension            []Extension                      `json:"modifierExtension,omitempty"`
	Url                          *string                          `json:"url,omitempty"`
	Identifier                   []Identifier                     `json:"identifier,omitempty"`
	Version                      *string                          `json:"version,omitempty"`
	Name                         *string                          `json:"name,omitempty"`
	Title                        *string                          `json:"title,omitempty"`
	Subtitle                     *string                          `json:"subtitle,omitempty"`
	Status                       PublicationStatus                `json:"status"`
	Experimental                 *bool                            `json:"experimental,omitempty"`
	SubjectCodeableConcept       *CodeableConcept                 `json:"subjectCodeableConcept,omitempty"`
	SubjectReference             *Reference                       `json:"subjectReference,omitempty"`
	Date                         *string                          `json:"date,omitempty"`
	Publisher                    *string                          `json:"publisher,omitempty"`
	Contact                      []ContactDetail                  `json:"contact,omitempty"`
	Description                  *string                          `json:"description,omitempty"`
	UseContext                   []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                 []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                      *string                          `json:"purpose,omitempty"`
	Usage                        *string                          `json:"usage,omitempty"`
	ApprovalDate                 *string                          `json:"approvalDate,omitempty"`
	LastReviewDate               *string                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod              *Period                          `json:"effectivePeriod,omitempty"`
	Topic                        []CodeableConcept                `json:"topic,omitempty"`
	Author                       []ContactDetail                  `json:"author,omitempty"`
	Editor                       []ContactDetail                  `json:"editor,omitempty"`
	Reviewer                     []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser                     []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact              []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Library                      []string                         `json:"library,omitempty"`
	Kind                         *RequestResourceType             `json:"kind,omitempty"`
	Profile                      *string                          `json:"profile,omitempty"`
	Code                         *CodeableConcept                 `json:"code,omitempty"`
	Intent                       *RequestIntent                   `json:"intent,omitempty"`
	Priority                     *RequestPriority                 `json:"priority,omitempty"`
	DoNotPerform                 *bool                            `json:"doNotPerform,omitempty"`
	TimingTiming                 *Timing                          `json:"timingTiming,omitempty"`
	TimingDateTime               *string                          `json:"timingDateTime,omitempty"`
	TimingAge                    *Age                             `json:"timingAge,omitempty"`
	TimingPeriod                 *Period                          `json:"timingPeriod,omitempty"`
	TimingRange                  *Range                           `json:"timingRange,omitempty"`
	TimingDuration               *Duration                        `json:"timingDuration,omitempty"`
	Location                     *Reference                       `json:"location,omitempty"`
	Participant                  []ActivityDefinitionParticipant  `json:"participant,omitempty"`
	ProductReference             *Reference                       `json:"productReference,omitempty"`
	ProductCodeableConcept       *CodeableConcept                 `json:"productCodeableConcept,omitempty"`
	Quantity                     *Quantity                        `json:"quantity,omitempty"`
	Dosage                       []Dosage                         `json:"dosage,omitempty"`
	BodySite                     []CodeableConcept                `json:"bodySite,omitempty"`
	SpecimenRequirement          []Reference                      `json:"specimenRequirement,omitempty"`
	ObservationRequirement       []Reference                      `json:"observationRequirement,omitempty"`
	ObservationResultRequirement []Reference                      `json:"observationResultRequirement,omitempty"`
	Transform                    *string                          `json:"transform,omitempty"`
	DynamicValue                 []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`
}

// Indicates who should participate in performing the action described.
type ActivityDefinitionParticipant struct {
	ID                *string               `json:"id,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Type              ActionParticipantType `json:"type"`
	Role              *CodeableConcept      `json:"role,omitempty"`
}

// Dynamic values that will be evaluated to produce values for elements of the resulting resource. For example, if the dosage of a medication must be computed based on the patient's weight, a dynamic value would be used to specify an expression that calculated the weight, and the path on the request resource that would contain the result.
// Dynamic values are applied in the order in which they are defined in the ActivityDefinition. Note that if both a transform and dynamic values are specified, the dynamic values will be applied to the result of the transform.
type ActivityDefinitionDynamicValue struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Path              string      `json:"path"`
	Expression        Expression  `json:"expression"`
}

// This function returns resource reference information
func (r ActivityDefinition) ResourceRef() (string, *string) {
	return "ActivityDefinition", r.ID
}

// This function returns resource reference information
func (r ActivityDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherActivityDefinition ActivityDefinition

// MarshalJSON marshals the given ActivityDefinition as JSON into a byte slice
func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActivityDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActivityDefinition: OtherActivityDefinition(r),
		ResourceType:            "ActivityDefinition",
	})
}

// UnmarshalActivityDefinition unmarshals a ActivityDefinition.
func UnmarshalActivityDefinition(b []byte) (ActivityDefinition, error) {
	var activityDefinition ActivityDefinition
	if err := json.Unmarshal(b, &activityDefinition); err != nil {
		return activityDefinition, err
	}
	return activityDefinition, nil
}
