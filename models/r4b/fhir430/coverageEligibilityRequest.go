package fhir430

import "encoding/json"

// CoverageEligibilityRequest is documented here http://hl7.org/fhir/StructureDefinition/CoverageEligibilityRequest
// The CoverageEligibilityRequest provides patient and insurance coverage information to an insurer for them to respond, in the form of an CoverageEligibilityResponse, with information regarding whether the stated coverage is valid and in-force and optionally to provide the insurance details of the policy.
type CoverageEligibilityRequest struct {
	ID                *string                                    `json:"id,omitempty"`
	Meta              *Meta                                      `json:"meta,omitempty"`
	ImplicitRules     *string                                    `json:"implicitRules,omitempty"`
	Language          *string                                    `json:"language,omitempty"`
	Text              *Narrative                                 `json:"text,omitempty"`
	Contained         []json.RawMessage                          `json:"contained,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes               `json:"status"`
	Priority          *CodeableConcept                           `json:"priority,omitempty"`
	Purpose           []EligibilityRequestPurpose                `json:"purpose"`
	Patient           Reference                                  `json:"patient"`
	ServicedDate      *string                                    `json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                    `json:"servicedPeriod,omitempty"`
	Created           string                                     `json:"created"`
	Enterer           *Reference                                 `json:"enterer,omitempty"`
	Provider          *Reference                                 `json:"provider,omitempty"`
	Insurer           Reference                                  `json:"insurer"`
	Facility          *Reference                                 `json:"facility,omitempty"`
	SupportingInfo    []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty"`
	Insurance         []CoverageEligibilityRequestInsurance      `json:"insurance,omitempty"`
	Item              []CoverageEligibilityRequestItem           `json:"item,omitempty"`
}

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
// Often there are multiple jurisdiction specific valuesets which are required.
type CoverageEligibilityRequestSupportingInfo struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Sequence          int         `json:"sequence"`
	Information       Reference   `json:"information"`
	AppliesToAll      *bool       `json:"appliesToAll,omitempty"`
}

// Financial instruments for reimbursement for the health care products and services.
// All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
type CoverageEligibilityRequestInsurance struct {
	ID                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Focal               *bool       `json:"focal,omitempty"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
}

// Service categories or billable services for which benefit details and/or an authorization prior to service delivery may be required by the payor.
type CoverageEligibilityRequestItem struct {
	ID                     *string                                   `json:"id,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	SupportingInfoSequence []int                                     `json:"supportingInfoSequence,omitempty"`
	Category               *CodeableConcept                          `json:"category,omitempty"`
	ProductOrService       *CodeableConcept                          `json:"productOrService,omitempty"`
	Modifier               []CodeableConcept                         `json:"modifier,omitempty"`
	Provider               *Reference                                `json:"provider,omitempty"`
	Quantity               *Quantity                                 `json:"quantity,omitempty"`
	UnitPrice              *Money                                    `json:"unitPrice,omitempty"`
	Facility               *Reference                                `json:"facility,omitempty"`
	Diagnosis              []CoverageEligibilityRequestItemDiagnosis `json:"diagnosis,omitempty"`
	Detail                 []Reference                               `json:"detail,omitempty"`
}

// Patient diagnosis for which care is sought.
type CoverageEligibilityRequestItemDiagnosis struct {
	ID                       *string          `json:"id,omitempty"`
	Extension                []Extension      `json:"extension,omitempty"`
	ModifierExtension        []Extension      `json:"modifierExtension,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference       `json:"diagnosisReference,omitempty"`
}

// This function returns resource reference information
func (r CoverageEligibilityRequest) ResourceRef() (string, *string) {
	return "CoverageEligibilityRequest", r.ID
}

// This function returns resource reference information
func (r CoverageEligibilityRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCoverageEligibilityRequest CoverageEligibilityRequest

// MarshalJSON marshals the given CoverageEligibilityRequest as JSON into a byte slice
func (r CoverageEligibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityRequest: OtherCoverageEligibilityRequest(r),
		ResourceType:                    "CoverageEligibilityRequest",
	})
}

// UnmarshalCoverageEligibilityRequest unmarshals a CoverageEligibilityRequest.
func UnmarshalCoverageEligibilityRequest(b []byte) (CoverageEligibilityRequest, error) {
	var coverageEligibilityRequest CoverageEligibilityRequest
	if err := json.Unmarshal(b, &coverageEligibilityRequest); err != nil {
		return coverageEligibilityRequest, err
	}
	return coverageEligibilityRequest, nil
}
