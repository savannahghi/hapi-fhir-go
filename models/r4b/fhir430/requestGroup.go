package fhir430

import "encoding/json"

// RequestGroup is documented here http://hl7.org/fhir/StructureDefinition/RequestGroup
// A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup struct {
	ID                    *string              `json:"id,omitempty"`
	Meta                  *Meta                `json:"meta,omitempty"`
	ImplicitRules         *string              `json:"implicitRules,omitempty"`
	Language              *string              `json:"language,omitempty"`
	Text                  *Narrative           `json:"text,omitempty"`
	Contained             []json.RawMessage    `json:"contained,omitempty"`
	Extension             []Extension          `json:"extension,omitempty"`
	ModifierExtension     []Extension          `json:"modifierExtension,omitempty"`
	Identifier            []Identifier         `json:"identifier,omitempty"`
	InstantiatesCanonical []string             `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string             `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference          `json:"basedOn,omitempty"`
	Replaces              []Reference          `json:"replaces,omitempty"`
	GroupIdentifier       *Identifier          `json:"groupIdentifier,omitempty"`
	Status                RequestStatus        `json:"status"`
	Intent                RequestIntent        `json:"intent"`
	Priority              *RequestPriority     `json:"priority,omitempty"`
	Code                  *CodeableConcept     `json:"code,omitempty"`
	Subject               *Reference           `json:"subject,omitempty"`
	Encounter             *Reference           `json:"encounter,omitempty"`
	AuthoredOn            *string              `json:"authoredOn,omitempty"`
	Author                *Reference           `json:"author,omitempty"`
	ReasonCode            []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference       []Reference          `json:"reasonReference,omitempty"`
	Note                  []Annotation         `json:"note,omitempty"`
	Action                []RequestGroupAction `json:"action,omitempty"`
}

// The actions, if any, produced by the evaluation of the artifact.
type RequestGroupAction struct {
	ID                  *string                           `json:"id,omitempty"`
	Extension           []Extension                       `json:"extension,omitempty"`
	ModifierExtension   []Extension                       `json:"modifierExtension,omitempty"`
	Prefix              *string                           `json:"prefix,omitempty"`
	Title               *string                           `json:"title,omitempty"`
	Description         *string                           `json:"description,omitempty"`
	TextEquivalent      *string                           `json:"textEquivalent,omitempty"`
	Priority            *RequestPriority                  `json:"priority,omitempty"`
	Code                []CodeableConcept                 `json:"code,omitempty"`
	Documentation       []RelatedArtifact                 `json:"documentation,omitempty"`
	Condition           []RequestGroupActionCondition     `json:"condition,omitempty"`
	RelatedAction       []RequestGroupActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime      *string                           `json:"timingDateTime,omitempty"`
	TimingAge           *Age                              `json:"timingAge,omitempty"`
	TimingPeriod        *Period                           `json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                         `json:"timingDuration,omitempty"`
	TimingRange         *Range                            `json:"timingRange,omitempty"`
	TimingTiming        *Timing                           `json:"timingTiming,omitempty"`
	Participant         []Reference                       `json:"participant,omitempty"`
	Type                *CodeableConcept                  `json:"type,omitempty"`
	GroupingBehavior    *ActionGroupingBehavior           `json:"groupingBehavior,omitempty"`
	SelectionBehavior   *ActionSelectionBehavior          `json:"selectionBehavior,omitempty"`
	RequiredBehavior    *ActionRequiredBehavior           `json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *ActionPrecheckBehavior           `json:"precheckBehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior        `json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                        `json:"resource,omitempty"`
	Action              []RequestGroupAction              `json:"action,omitempty"`
}

// An expression that describes applicability criteria, or start/stop conditions for the action.
// When multiple conditions of the same kind are present, the effects are combined using AND semantics, so the overall condition is true only if all of the conditions are true.
type RequestGroupActionCondition struct {
	ID                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Kind              ActionConditionKind `json:"kind"`
	Expression        *Expression         `json:"expression,omitempty"`
}

// A relationship to another action such as "before" or "30-60 minutes after start of".
type RequestGroupActionRelatedAction struct {
	ID                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	ActionId          string                 `json:"actionId"`
	Relationship      ActionRelationshipType `json:"relationship"`
	OffsetDuration    *Duration              `json:"offsetDuration,omitempty"`
	OffsetRange       *Range                 `json:"offsetRange,omitempty"`
}

// This function returns resource reference information
func (r RequestGroup) ResourceRef() (string, *string) {
	return "RequestGroup", r.ID
}

// This function returns resource reference information
func (r RequestGroup) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherRequestGroup RequestGroup

// MarshalJSON marshals the given RequestGroup as JSON into a byte slice
func (r RequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestGroup: OtherRequestGroup(r),
		ResourceType:      "RequestGroup",
	})
}

// UnmarshalRequestGroup unmarshals a RequestGroup.
func UnmarshalRequestGroup(b []byte) (RequestGroup, error) {
	var requestGroup RequestGroup
	if err := json.Unmarshal(b, &requestGroup); err != nil {
		return requestGroup, err
	}
	return requestGroup, nil
}
