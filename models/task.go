package models

import (
	"encoding/json"
	"errors"

	"github.com/savannahghi/scalarutils"
)

// FHIRTask is a FHIR task resource data class.
type FHIRTask struct {
	ID                    *string                `json:"id,omitempty"`
	Meta                  *FHIRMeta              `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *FHIRNarrative         `json:"text,omitempty"`
	Extension             []*FHIRExtension       `json:"extension,omitempty"`
	ModifierExtension     []*FHIRExtension       `json:"modifierExtension,omitempty"`
	Identifier            []*FHIRIdentifier      `json:"identifier,omitempty"`
	InstantiatesCanonical *scalarutils.Canonical `json:"instantiatesCanonical,omitempty"`
	InstantiatesURI       *scalarutils.URI       `json:"instantiatesUri,omitempty"`
	BasedOn               []*FHIRReference       `json:"basedOn,omitempty"`
	GroupIdentifier       *FHIRIdentifier        `json:"groupIdentifier,omitempty"`
	PartOf                []*FHIRReference       `json:"partOf,omitempty"`
	Status                *scalarutils.Code      `json:"status"`
	StatusReason          *FHIRCodeableConcept   `json:"statusReason,omitempty"`
	BusinessStatus        *FHIRCodeableConcept   `json:"businessStatus,omitempty"`
	Intent                string                 `json:"intent"`
	Priority              *scalarutils.Code      `json:"priority,omitempty"`
	Code                  *FHIRCodeableConcept   `json:"code,omitempty"`
	Description           string                 `json:"description,omitempty"`
	Focus                 *FHIRReference         `json:"focus,omitempty"`
	For                   *FHIRReference         `json:"for,omitempty"`
	Encounter             *FHIRReference         `json:"encounter,omitempty"`
	ExecutionPeriod       *FHIRPeriod            `json:"executionPeriod,omitempty"`
	AuthoredOn            *string                `json:"authoredOn,omitempty"`
	LastModified          *string                `json:"lastModified,omitempty"`
	Requester             *FHIRReference         `json:"requester,omitempty"`
	PerformerType         []*FHIRCodeableConcept `json:"performerType,omitempty"`
	Owner                 *FHIRReference         `json:"owner,omitempty"`
	Location              *FHIRReference         `json:"location,omitempty"`
	ReasonCode            *FHIRCodeableConcept   `json:"reasonCode,omitempty"`
	ReasonReference       *FHIRReference         `json:"reasonReference,omitempty"`
	Insurance             []*FHIRReference       `json:"insurance,omitempty"`
	Note                  []*FHIRAnnotation      `json:"note,omitempty"`
	RelevantHistory       []*FHIRReference       `json:"relevantHistory,omitempty"`
	Restriction           *TaskRestriction       `json:"restriction,omitempty"`
	Input                 []TaskInput            `json:"input,omitempty"`
	Output                []TaskOutput           `json:"output,omitempty"`
}

// GetServiceRequestIDFromTask is used to extract the referral (service request) ID from a task.
func (t *FHIRTask) GetServiceRequestIDFromTask() (string, error) {
	if t == nil {
		return "", errors.New("task is nil")
	}

	var referralID string

	for _, serviceRequest := range t.BasedOn {
		if serviceRequest.Type != nil && string(*serviceRequest.Type) == ReferralServiceRequestType.String() {
			referralID = "ServiceRequest/" + *serviceRequest.ID

			break
		}
	}

	return referralID, nil
}

// TaskRestriction models the constraints on fulfillment tasks.
type TaskRestriction struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []*FHIRExtension `json:"extension,omitempty"`
	ModifierExtension []*FHIRExtension `json:"modifierExtension,omitempty"`
	Repetitions       *int             `json:"repetitions,omitempty"`
	Period            *FHIRPeriod      `json:"period,omitempty"`
	Recipient         []*FHIRReference `json:"recipient,omitempty"`
}

// TaskInput models the information needed to perform a task.
type TaskInput struct {
	ID                *string              `json:"id,omitempty"`
	Extension         []*FHIRExtension     `json:"extension,omitempty"`
	ModifierExtension []*FHIRExtension     `json:"modifierExtension,omitempty"`
	Type              *FHIRCodeableConcept `json:"type"`
	ValueBase64Binary string               `json:"valueBase64Binary"`
	ValueBoolean      bool                 `json:"valueBoolean"`
	ValueCanonical    string               `json:"valueCanonical"`
	ValueCode         string               `json:"valueCode"`
	ValueDate         string               `json:"valueDate"`
	ValueDateTime     string               `json:"valueDateTime"`
	ValueDecimal      json.Number          `json:"valueDecimal"`
	ValueID           string               `json:"valueId"`
	ValueInstant      string               `json:"valueInstant"`
	ValueInteger      int                  `json:"valueInteger"`
	ValueMarkdown     string               `json:"valueMarkdown"`
	ValueOID          string               `json:"valueOid"`
	ValuePositiveInt  int                  `json:"valuePositiveInt"`
	ValueString       string               `json:"valueString"`
	ValueTime         string               `json:"valueTime"`
	ValueUnsignedInt  int                  `json:"valueUnsignedInt"`
	ValueURI          string               `json:"valueUri"`
	ValueURL          string               `json:"valueUrl"`
}

// TaskOutput models the information produced as part of task.
type TaskOutput struct {
	ID                *string              `json:"id,omitempty"`
	Extension         []*FHIRExtension     `json:"extension,omitempty"`
	ModifierExtension []*FHIRExtension     `json:"modifierExtension,omitempty"`
	Type              *FHIRCodeableConcept `json:"type"`
	ValueBase64Binary string               `json:"valueBase64Binary"`
	ValueBoolean      bool                 `json:"valueBoolean"`
	ValueCanonical    string               `json:"valueCanonical"`
	ValueCode         string               `json:"valueCode"`
	ValueDate         string               `json:"valueDate"`
	ValueDateTime     string               `json:"valueDateTime"`
	ValueDecimal      json.Number          `json:"valueDecimal"`
	ValueID           string               `json:"valueId"`
	ValueInstant      string               `json:"valueInstant"`
	ValueInteger      int                  `json:"valueInteger"`
	ValueMarkdown     string               `json:"valueMarkdown"`
	ValueOID          string               `json:"valueOid"`
	ValuePositiveInt  int                  `json:"valuePositiveInt"`
	ValueString       string               `json:"valueString"`
	ValueTime         string               `json:"valueTime"`
	ValueUnsignedInt  int                  `json:"valueUnsignedInt"`
	ValueURI          string               `json:"valueUri"`
	ValueURL          string               `json:"valueUrl"`
	ValueUUUID        string               `json:"valueUuid"`
}

// FHIRTaskRelayPayload is used to return single instances of Task.
type FHIRTaskRelayPayload struct {
	Resource *FHIRTask `json:"resource,omitempty"`
}
