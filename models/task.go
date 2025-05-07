package models

import (
	"encoding/json"
	"errors"

	"github.com/savannahghi/scalarutils"
)

// Task is a  task resource data class.
type Task struct {
	ID                    *string                `json:"id,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Extension             []*Extension           `json:"extension,omitempty"`
	ModifierExtension     []*Extension           `json:"modifierExtension,omitempty"`
	Identifier            []*Identifier          `json:"identifier,omitempty"`
	InstantiatesCanonical *scalarutils.Canonical `json:"instantiatesCanonical,omitempty"`
	InstantiatesURI       *string                `json:"instantiatesUri,omitempty"`
	BasedOn               []*Reference           `json:"basedOn,omitempty"`
	GroupIdentifier       *Identifier            `json:"groupIdentifier,omitempty"`
	PartOf                []*Reference           `json:"partOf,omitempty"`
	Status                *string                `json:"status"`
	StatusReason          *CodeableConcept       `json:"statusReason,omitempty"`
	BusinessStatus        *CodeableConcept       `json:"businessStatus,omitempty"`
	Intent                string                 `json:"intent"`
	Priority              *string                `json:"priority,omitempty"`
	Code                  *CodeableConcept       `json:"code,omitempty"`
	Description           string                 `json:"description,omitempty"`
	Focus                 *Reference             `json:"focus,omitempty"`
	For                   *Reference             `json:"for,omitempty"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	ExecutionPeriod       *Period                `json:"executionPeriod,omitempty"`
	AuthoredOn            *string                `json:"authoredOn,omitempty"`
	LastModified          *string                `json:"lastModified,omitempty"`
	Requester             *Reference             `json:"requester,omitempty"`
	PerformerType         []*CodeableConcept     `json:"performerType,omitempty"`
	Owner                 *Reference             `json:"owner,omitempty"`
	Location              *Reference             `json:"location,omitempty"`
	ReasonCode            *CodeableConcept       `json:"reasonCode,omitempty"`
	ReasonReference       *Reference             `json:"reasonReference,omitempty"`
	Insurance             []*Reference           `json:"insurance,omitempty"`
	Note                  []*Annotation          `json:"note,omitempty"`
	RelevantHistory       []*Reference           `json:"relevantHistory,omitempty"`
	Restriction           *TaskRestriction       `json:"restriction,omitempty"`
	Input                 []TaskInput            `json:"input,omitempty"`
	Output                []TaskOutput           `json:"output,omitempty"`
}

// GetServiceRequestIDFromTask is used to extract the referral (service request) ID from a task.
func (t *Task) GetServiceRequestIDFromTask() (string, error) {
	if t == nil {
		return "", errors.New("task is nil")
	}

	var referralID string

	for _, serviceRequest := range t.BasedOn {
		if serviceRequest.Type != nil && *serviceRequest.Type == ReferralServiceRequestType.String() {
			referralID = "ServiceRequest/" + *serviceRequest.ID

			break
		}
	}

	return referralID, nil
}

// TaskRestriction models the constraints on fulfillment tasks.
type TaskRestriction struct {
	ID                *string      `json:"id,omitempty"`
	Extension         []*Extension `json:"extension,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Repetitions       *int         `json:"repetitions,omitempty"`
	Period            *Period      `json:"period,omitempty"`
	Recipient         []*Reference `json:"recipient,omitempty"`
}

// TaskInput models the information needed to perform a task.
type TaskInput struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []*Extension     `json:"extension,omitempty"`
	ModifierExtension []*Extension     `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type"`
	ValueBase64Binary string           `json:"valueBase64Binary"`
	ValueBoolean      bool             `json:"valueBoolean"`
	ValueCanonical    string           `json:"valueCanonical"`
	ValueCode         string           `json:"valueCode"`
	ValueDate         string           `json:"valueDate"`
	ValueDateTime     string           `json:"valueDateTime"`
	ValueDecimal      json.Number      `json:"valueDecimal"`
	ValueID           string           `json:"valueId"`
	ValueInstant      string           `json:"valueInstant"`
	ValueInteger      int              `json:"valueInteger"`
	ValueMarkdown     string           `json:"valueMarkdown"`
	ValueOID          string           `json:"valueOid"`
	ValuePositiveInt  int              `json:"valuePositiveInt"`
	ValueString       string           `json:"valueString"`
	ValueTime         string           `json:"valueTime"`
	ValueUnsignedInt  int              `json:"valueUnsignedInt"`
	ValueURI          string           `json:"valueUri"`
	ValueURL          string           `json:"valueUrl"`
}

// TaskOutput models the information produced as part of task.
type TaskOutput struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []*Extension     `json:"extension,omitempty"`
	ModifierExtension []*Extension     `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type"`
	ValueBase64Binary string           `json:"valueBase64Binary"`
	ValueBoolean      bool             `json:"valueBoolean"`
	ValueCanonical    string           `json:"valueCanonical"`
	ValueCode         string           `json:"valueCode"`
	ValueDate         string           `json:"valueDate"`
	ValueDateTime     string           `json:"valueDateTime"`
	ValueDecimal      json.Number      `json:"valueDecimal"`
	ValueID           string           `json:"valueId"`
	ValueInstant      string           `json:"valueInstant"`
	ValueInteger      int              `json:"valueInteger"`
	ValueMarkdown     string           `json:"valueMarkdown"`
	ValueOID          string           `json:"valueOid"`
	ValuePositiveInt  int              `json:"valuePositiveInt"`
	ValueString       string           `json:"valueString"`
	ValueTime         string           `json:"valueTime"`
	ValueUnsignedInt  int              `json:"valueUnsignedInt"`
	ValueURI          string           `json:"valueUri"`
	ValueURL          string           `json:"valueUrl"`
	ValueUUUID        string           `json:"valueUuid"`
}

// TaskRelayPayload is used to return single instances of Task.
type TaskRelayPayload struct {
	Resource *Task `json:"resource,omitempty"`
}
